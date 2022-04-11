package controller

import (
	"AlthCart/models"
	"github.com/kataras/iris/v12"
)

type IndexData struct {
	Products  []models.Products
	CartCount int64
}

func IndexHandler(ctx iris.Context) {
	uId := models.UserIdJWT(ctx.GetCookie("token"))
	dataProduct, err := models.GetProducts()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get Product List").DetailErr(err))
		return
	}

	totalCart, err := models.GetUserQuantityCart(uId)
	if err != nil {
		totalCart = 0
	}

	IndexData := IndexData{Products: dataProduct, CartCount: totalCart}
	ctx.ViewData("Data", IndexData)
	ctx.View("index.html")
}

func LoginHandler(ctx iris.Context) {
	ctx.View("login.html")
}
