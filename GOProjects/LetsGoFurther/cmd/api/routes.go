package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/artifacts", app.listArtifactsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/artifacts", app.createArtifactHandler)
	router.HandlerFunc(http.MethodGet, "/v1/artifacts/:id", app.showArtifactHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/artifacts/:id", app.updateArtifactHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/artifacts/:id", app.deleteArtifactHandler)
	return app.recoverPanic(app.rateLimit(router))
}
