package app

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/navythenerd/lionrouter"
)

func (app *Application) NotFoundHandler() http.Handler {
	files := []string{
		fmt.Sprintf("%s/%s", app.config.Views, "master.gohtml"),
		fmt.Sprintf("%s/%s", app.config.Views, "header.gohtml"),
		fmt.Sprintf("%s/%s", app.config.Views, "404.gohtml"),
	}
	tmpl := template.Must(template.ParseFiles(files...))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		data := ViewData{
			Title:       app.config.Title,
			Description: app.config.Description,
			Website:     app.config.Website,
			Nametag:     app.config.Nametag,
			Avatar:      app.config.Avatar,
		}

		_ = tmpl.Execute(w, data)
	})
}

func (app *Application) treeHandler() http.Handler {
	files := []string{
		fmt.Sprintf("%s/%s", app.config.Views, "master.gohtml"),
		fmt.Sprintf("%s/%s", app.config.Views, "header.gohtml"),
		fmt.Sprintf("%s/%s", app.config.Views, "links.gohtml"),
	}
	tmpl := template.Must(template.ParseFiles(files...))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:       app.config.Title,
			Description: app.config.Description,
			Links:       app.config.Links,
			Website:     app.config.Website,
			Nametag:     app.config.Nametag,
			Avatar:      app.config.Avatar,
		}

		_ = tmpl.Execute(w, data)
	})
}

func (app *Application) staticHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := lionrouter.Param(r.Context(), "file")
		path := fmt.Sprintf("%s/%s", app.config.Assets, file)

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			app.NotFoundHandler().ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, fmt.Sprintf("%s/%s", app.config.Assets, file))
	})
}
