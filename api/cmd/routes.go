package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/adventure", app.adventure)
	router.HandlerFunc(http.MethodPost, "/adventure", app.adventurePost)
	return router
}
