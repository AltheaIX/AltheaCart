package controller

import (
	"AlthCart/models"
	"github.com/kataras/iris/v12"
)

type CartData struct {
	Carts      []models.Carts
	CartCount  int64
	TotalPrice int64
}

func CartHandler(ctx iris.Context) {
	uId := models.UserIdJWT(ctx.GetCookie("token"))
	userCart, err := models.GetUserCart(uId)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get user's cart.").DetailErr(err))
		return
	}

	quantityCart, err := models.GetUserQuantityCart(uId)
	if err != nil {
		quantityCart = 0
	}

	totalCart := models.GetUserTotal(userCart)

	cartData := CartData{Carts: userCart, CartCount: quantityCart, TotalPrice: totalCart}

	ctx.ViewData("Data", cartData)
	ctx.View("cart.html")
}

func CartsTotal(ctx iris.Context) {
	uId := models.UserIdJWT(ctx.GetCookie("token"))
	userCart, err := models.GetUserCart(uId)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to get userCart"))
		return
	}

	totalCart := models.GetUserTotal(userCart)
	ctx.JSON(map[string]interface{}{
		"status": iris.StatusOK,
		"total":  totalCart,
	})
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
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title("Error when trying to remove carts").DetailErr(err))
		return
	}
	ctx.JSON(map[string]interface{}{
		"status":  iris.StatusOK,
		"message": "Remove Success.",
	})
}
