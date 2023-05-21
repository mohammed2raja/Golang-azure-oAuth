package handlers

import (
	"auth-server/read-mail/internal/db/repositories"
	"auth-server/read-mail/internal/models"
	"auth-server/read-mail/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

func ReadProfileHandler(w http.ResponseWriter, r *http.Request) {
	// read user name from query
	queryParams := r.URL.Query()
	username := queryParams.Get("name")

	// Create a new instance of UserRepository
	userRepo, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	user, err := userRepo.GetByName(username)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := services.CallProfileAPI(user.Token)
	if err != nil {
		log.Println("Error calling MsGraph API:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Data: string(responseBody),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshaling JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
