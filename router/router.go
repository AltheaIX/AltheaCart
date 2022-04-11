package router

import (
	"AlthCart/controller"
	"github.com/kataras/iris/v12"
)

func Router() {
	app := iris.New()

	app.HandleDir("/assets", iris.Dir("./controller/assets"))
	tmpl := iris.HTML("./controller/templates", ".html")
	tmpl.Reload(true)

	main := app.Party("/")
	{
		app.RegisterView(tmpl)
		app.Get("/login", controller.LoginHandler)
		main.Use(controller.CheckAuth)
		main.Get("/", controller.IndexHandler)
		main.Get("/carts", controller.CartHandler)
		main.Get("/logout", controller.LogoutHandler)
	}

	api := app.Party("/api")
	{
		api.Use(iris.Compression)
		api.Post("/cart/add", controller.CheckAuth, controller.CartsAdd)
		api.Post("/cart/remove", controller.CheckAuth, controller.CartsRemove)
		api.Post("/login", controller.AuthHandler)
		api.Get("/product", controller.GetProductList)
		api.Get("/product/{id:int64}", controller.GetProduct)
	}
	app.Listen(":8003")
}
