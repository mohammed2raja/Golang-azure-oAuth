package main

import (
	"auth-server/read-mail/web/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// Setup API routes
	handler := routes.SetupAPIRoutes()

	// Start the HTTP server
	log.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", handler))

}
