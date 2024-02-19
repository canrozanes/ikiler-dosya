package router

import (
	"net/http"

	"ikiler-dosya/frontend"
	"ikiler-dosya/pkg/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Router(serverOrigin, audience, auth0Domain string) http.Handler {

	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)

	router.HandleFunc("/api/messages/public", middleware.PublicApiHandler)
	router.Handle("/api/messages/protected", middleware.ValidateJWT(audience, auth0Domain, http.HandlerFunc(middleware.ProtectedApiHandler)))
	router.Handle("/api/messages/admin",
		middleware.ValidateJWT(audience, auth0Domain,
			middleware.ValidatePermissions([]string{"read:admin-messages"},
				http.HandlerFunc(middleware.AdminApiHandler))))

	spaHandler := frontend.CreateSpaHandler(serverOrigin, auth0Domain)
	router.NotFound(spaHandler)

	return middleware.HandleCacheControl(router)
}
