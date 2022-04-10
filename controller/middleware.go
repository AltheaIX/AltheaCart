package controller

import (
	"AlthCart/models"
	"github.com/kataras/iris/v12"
)

func CheckAuth(ctx iris.Context) {
	token := ctx.GetCookie("token")
	_, err := models.VerifyJWT(token)
	if err != nil {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}

	ctx.Next()
}
