package handlers

import (
	"auth-server/read-mail/internal/db/repositories"
	"auth-server/read-mail/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	AccessToken string `json:"access_token,omitempty"`
	User        string `json:"user,omitempty"`
	Email       string `json:"email,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

func ReadAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var result Result
	json.Unmarshal(body, &result)

	// Create a new instance of UserRepository
	userRepo, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Use the repository methods
	user := &models.User{
		Name:  result.User,
		Token: result.AccessToken,
	}

	err = userRepo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record inserted/updated successfully!")

}
