package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
	"golang.org/x/crypto/bcrypt"
)

// Helper function that creates a new session token and sets a browser cookie with that token
func createCookie(w http.ResponseWriter, r *http.Request, id string) (http.ResponseWriter, string, error) {
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		http.Error(w, "Failed to create token", http.StatusBadRequest)
		return w, "", err
	}

	// send it back

	// when hosting website change secure bool to TRUE
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return w, tokenString, err
}

// Used this tutorial: https://mj-go.in/golang/rest-api-with-sql-db-in-go

// POST function used to create a user. Takes in username and password strings in json, hashes the password, and stores data in database. A cookie is also made for automatic login, and the newly created user is sent back to the front end.
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
	var t string
	// create session token and cookie
	w, t, err = createCookie(w, r, body.Username)
	if err != nil {
		return
	}
	// Create the user
	user := models.User{Username: body.Username, Password: string(hash), SessionToken: t}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		http.Error(w, "Username already taken", http.StatusInternalServerError)
		return
	}

	// Create income amount of 0 for new user
	income := models.Income{Username: body.Username, Amount: 0}

	result = initializers.DB.Create(&income) // pass pointer of data to Create

	if result.Error != nil {
		http.Error(w, "Username already taken", http.StatusInternalServerError)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Signed up\n"))
}

// POST function to enter website with existing user. checks to see if input username exists, and if input password matches associated password. if valid, creates cookie for login session. user data is returned to the front end.
func Login(w http.ResponseWriter, r *http.Request) {
	// Get the email and pass off req body
	var body struct {
		Username string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	// Compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}
	var token string
	// Create a session token and cookie that stores the token
	w, token, err = createCookie(w, r, user.Username)
	if err != nil {
		return
	}

	// Update session token in DB with new token
	user.SessionToken = token
	initializers.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged In\n"))

}

func InputIncome(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string
		Amount   uint
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.Income
	initializers.DB.First(&user, "username = ?", body.Username)
	user.Amount = body.Amount
	initializers.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(body.Amount)))

}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username    string
		ExpenseType string
		ExpenseName string
		Amount      uint
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch body.ExpenseType {
	case "HomeUtils":
		exp := models.Home_Uts{Username: body.Username, ExpenseName: body.ExpenseName, Amount: body.Amount}

		result := initializers.DB.Create(&exp) // pass pointer of data to Create

		if result.Error != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}
	case "Transportation":
		exp := models.Trans{Username: body.Username, ExpenseName: body.ExpenseName, Amount: body.Amount}

		result := initializers.DB.Create(&exp) // pass pointer of data to Create

		if result.Error != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}
	case "Food":
		exp := models.Food{Username: body.Username, ExpenseName: body.ExpenseName, Amount: body.Amount}

		result := initializers.DB.Create(&exp) // pass pointer of data to Create

		if result.Error != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}
	case "Education":
		exp := models.Edu{Username: body.Username, ExpenseName: body.ExpenseName, Amount: body.Amount}

		result := initializers.DB.Create(&exp) // pass pointer of data to Create

		if result.Error != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}
	case "Health":
		exp := models.Health{Username: body.Username, ExpenseName: body.ExpenseName, Amount: body.Amount}

		result := initializers.DB.Create(&exp) // pass pointer of data to Create

		if result.Error != nil {
			http.Error(w, "Failed to create expense", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Failed to create expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Added\n"))
}
