package controller

import (
	"AlthCart/models"
	"github.com/kataras/iris/v12"
)

func CartHandler(ctx iris.Context) {
	token := ctx.GetCookie("token")
	uId := models.UserIdJWT(token)
	userCart, err := models.GetUserCart(uId)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get user's cart.").DetailErr(err))
		return
	}
	ctx.JSON(userCart)
}

func CartsAdd(ctx iris.Context) {
	uId := models.UserIdJWT(ctx.GetCookie("token"))
	id, err := ctx.PostValueInt64("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Parsing Error").DetailErr(err))
		return
	}

	err = models.AddCarts(id, uId)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to add carts").DetailErr(err))
		return
	}
	ctx.JSON(map[string]interface{}{
		"status":  iris.StatusOK,
		"message": "Added success.",
	})
}

func CartsRemove(ctx iris.Context) {
	uId := models.UserIdJWT(ctx.GetCookie("token"))
	id, err := ctx.PostValueInt64("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Parsing Error").DetailErr(err))
		return
	}

	err = models.RemoveCarts(id, uId)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to add carts").DetailErr(err))
		return
	}
	ctx.JSON(map[string]interface{}{
		"status":  iris.StatusOK,
		"message": "Remove Success.",
	})
}
