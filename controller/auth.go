package controller

import (
	"AlthCart/models"
	"database/sql"
	"github.com/kataras/iris/v12"
	"time"
)

func LoginHandler(ctx iris.Context) {
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	data, err := models.ValidateAuth(username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Writef("Wrong username or password!")
			return
		}
		ctx.StopWithError(iris.StatusInternalServerError, iris.NewProblem().Title("Auth Error").DetailErr(err))
	}
	token := models.CreateJWT(data)
	ctx.SetCookieKV("token", token, iris.CookieExpires(1*time.Hour))
	verify, err := models.VerifyJWT(token)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to validate JWT").DetailErr(err))
	}

	ctx.JSON(verify)
}

func LogoutHandler(ctx iris.Context) {
	ctx.RemoveCookie("token")
	ctx.Redirect("/login", iris.StatusTemporaryRedirect)
}
