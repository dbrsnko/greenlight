package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)
	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	return app.recoverPanic(router)
}
