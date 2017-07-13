package main

func main() {

	app = NewApp()
	app.Use(middleware1, middleware2)
	userController = app.Controller("/user")
	userController.Use(middleware1, middleware2)

	userController.POST("/", func(ctx *orange.Context) {

	}, middleware1, middlewar2)

	userController.GET("/", func(ctx *orange.Context) {

	}, middleware1, middlewar2)

	ns = app.NS("v1", path)
	ns.Use(middleware1, middleware2)
	userController = ns.Controller("users")
	userController.Use(middleware1, middleware2)
	userController.POST("/", func(ctx *orange.Context) {

	}, middleware1, middlewar2)

	userController.GET("/", func(ctx *orange.Context) {

	}, middleware1, middlewar2)

}
