package router

import (
	"AlthCart/controller"
	"github.com/kataras/iris/v12"
)

func Router() {
	app := iris.New()

	main := app.Party("/")
	{
		main.Use(controller.CheckAuth)
		main.Get("/", controller.GetProductList)
	}

	api := app.Party("/api")
	{
		api.Use(iris.Compression)
		api.Get("/login", controller.DoAuth)
		api.Get("/check", controller.CheckAuth)
		api.Get("/product", controller.GetProductList)
		api.Get("/product/{id:int64}", controller.GetProduct)
	}
	app.Listen(":8005")
}
