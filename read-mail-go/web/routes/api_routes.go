package routes

import (
	"auth-server/read-mail/internal/handlers"
	"net/http"

	"github.com/rs/cors"
)

func SetupAPIRoutes() http.Handler {

	mux := http.NewServeMux()

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)

	mux.HandleFunc("/", handlers.ReadAccessTokenHandler)

	mux.HandleFunc("/profile", handlers.ReadProfileHandler)
	mux.HandleFunc("/me/messages", handlers.ReadMessageHandler)
	return handler
}
