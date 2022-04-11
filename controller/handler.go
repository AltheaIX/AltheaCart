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
	IndexData := IndexData{Products: dataProduct, CartCount: totalCart}
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get Total Quantity Cart").DetailErr(err))
		return
	}
	ctx.ViewData("Data", IndexData)
	ctx.View("index.html")
}
