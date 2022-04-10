package controller

import (
	"AlthCart/models"
	"github.com/kataras/iris/v12"
)

func GetProductList(ctx iris.Context) {
	data, err := models.GetProducts()
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get list products.").DetailErr(err))
		return
	}
	ctx.JSON(data)
}

func GetProduct(ctx iris.Context) {
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Parameters Error").DetailErr(err))
	}
	data, err := models.GetProductById(id)
	if err != nil {
		ctx.JSON(err.Error())
		return
	}
	ctx.JSON(data)
}
