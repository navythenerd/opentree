package app

func (app *Application) registerRoutes() {
	app.router.NotFoundHandler = app.NotFoundHandler()
	app.router.Get("/static/*file", app.staticHandler())
	app.router.Get("/", app.treeHandler())
}
