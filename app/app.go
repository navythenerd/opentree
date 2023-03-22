package app

import (
	"net/http"

	"github.com/navythenerd/lionrouter"
)

type Application struct {
	config Config
	router *lionrouter.Router
}

func New(config Config) *Application {
	app := Application{
		config: config,
		router: lionrouter.New(),
	}

	app.registerRoutes()

	return &app
}

func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}
