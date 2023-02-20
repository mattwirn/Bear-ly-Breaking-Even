package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
	"golang.org/x/crypto/bcrypt"
)

// Used this tutorial: https://mj-go.in/golang/rest-api-with-sql-db-in-go

func Signup(w http.ResponseWriter, r *http.Request) {

	// Get the username/pass of req body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}

	// Create the user
	user := models.User{Username: body.Username, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		http.Error(w, "Username already taken", http.StatusInternalServerError)
		return
	}

	// Respond
	w.Write([]byte("Signed up\n"))
}

func Login(w http.ResponseWriter, r *http.Request) {

}
