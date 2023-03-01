package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Retrieve cookie from request
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusBadRequest)
		return
	}

	// Get token from value field
	token := cookie.Value

	// Look up user in database with token
	var user models.User
	initializers.DB.First(&user, "session_token = ?", token)

	if user.ID == 0 {
		http.Error(w, "Server side error", http.StatusInternalServerError)
		return
	}

	// Write username to json
	response, err := json.Marshal(user.Username)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Server side error", http.StatusInternalServerError)
		return
	}

	// Write json to response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
