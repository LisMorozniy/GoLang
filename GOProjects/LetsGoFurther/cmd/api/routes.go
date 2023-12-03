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
	router.HandlerFunc(http.MethodGet, "/v1/artifacts", app.requirePermission("artifacts:read", app.listArtifactsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/artifacts", app.requirePermission("artifacts:write", app.createArtifactHandler))
	router.HandlerFunc(http.MethodGet, "/v1/artifacts/:id", app.requirePermission("artifacts:read", app.showArtifactHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/artifacts/:id", app.requirePermission("artifacts:write", app.updateArtifactHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/artifacts/:id", app.requirePermission("artifacts:write", app.deleteArtifactHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))

}
