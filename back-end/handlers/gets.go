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
		http.Error(w, "ID error", http.StatusInternalServerError)
		return
	}

	// Interface to hold all write data
	allData := []interface{}{}

	// Write username to json
	username := user.Username
	//allData := getHomeUts(username)
	allData = append(allData, getUsername(username), getIncome(username), getTotalSpent(username), getAllExpenses(username))

	// Get income of user
	//allData = append(allData, getIncome(username))

	// Gather all expenses from DB
	//allData = append(allData, getAllExpenses(username))

	// Marshal all data to json
	response, err := json.Marshal(allData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Server side error", http.StatusInternalServerError)
		return
	}

	// Write json to response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func getUsername(username string) []interface{} {
	type username_ struct {
		Username string `json:"username"`
	}
	reformat := username_{Username: username}
	return []interface{}{reformat}
}

func getIncome(username string) []interface{} {
	type income struct {
		Income uint `json:"income"`
	}

	// Search table for income of user
	var inc models.Income
	initializers.DB.Find(&inc, "username = ?", username)

	// Reformat to front-end friendly struct
	reformat := income{Income: inc.Amount}

	// Add to interface and return
	return []interface{}{reformat}
}

func getTotalSpent(username string) []interface{} {
	type spent struct {
		TotalSpent uint `json:"totalspent"`
	}

	// Search table for user
	var user models.User
	initializers.DB.First(&user, "username = ?", username)

	// Add up to get total spent
	totalspent := user.HUTotal + user.FTotal + user.ETotal + user.HTotal + user.TTotal

	// Reformat to front-end friendly struct
	reformat := spent{TotalSpent: totalspent}

	// Add to interface and return
	return []interface{}{reformat}
}

// Template expense struct for all expenses to front-end
type expense struct {
	ExpenseType string `json:"expensetype"`
	ExpenseName string `json:"expensename"`
	Amount      uint   `json:"amount"`
}

// Template total expense category struct
type total struct {
	ExpenseType string `json:"expensetype"`
	Total       uint   `json:"total"`
}

func getHomeUts(username string) []interface{} {

	expenses := []interface{}{}

	// Search user table for expense total
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	total := total{
		ExpenseType: "HomeUtils",
		Total:       user.HUTotal,
	}
	expenses = append(expenses, total)

	// Search table for all expenses created the user
	var exps []models.Home_Uts
	result = initializers.DB.Find(&exps, "username = ?", username)

	if result.RowsAffected == 0 {
		return expenses
	}

	// Convert expense struct to new struct to be used by front-end
	for _, exp := range exps {
		reformat := expense{
			ExpenseType: "HomeUtils",
			ExpenseName: exp.ExpenseName,
			Amount:      exp.Amount,
		}
		expenses = append(expenses, reformat)
	}
	return expenses
}

func getTravel(username string) []interface{} {

	expenses := []interface{}{}

	// Search user table for expense total
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	total := total{
		ExpenseType: "Travel",
		Total:       user.TTotal,
	}
	expenses = append(expenses, total)

	// Search table for all expenses created the user
	var exps []models.Travel
	result = initializers.DB.Find(&exps, "username = ?", username)

	if result.RowsAffected == 0 {
		return expenses
	}

	// Convert expense struct to new struct to be used by front-end
	for _, exp := range exps {
		reformat := expense{
			ExpenseType: "Travel",
			ExpenseName: exp.ExpenseName,
			Amount:      exp.Amount,
		}
		expenses = append(expenses, reformat)
	}
	return expenses
}

func getFood(username string) []interface{} {

	expenses := []interface{}{}

	// Search user table for expense total
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	total := total{
		ExpenseType: "Food",
		Total:       user.FTotal,
	}
	expenses = append(expenses, total)

	// Search table for all expenses created the user
	var exps []models.Food
	result = initializers.DB.Find(&exps, "username = ?", username)

	if result.RowsAffected == 0 {
		return expenses
	}

	// Convert expense struct to new struct to be used by front-end
	for _, exp := range exps {
		reformat := expense{
			ExpenseType: "Food",
			ExpenseName: exp.ExpenseName,
			Amount:      exp.Amount,
		}
		expenses = append(expenses, reformat)
	}
	return expenses
}

func getEnt(username string) []interface{} {

	expenses := []interface{}{}

	// Search user table for expense total
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	total := total{
		ExpenseType: "Entertainment",
		Total:       user.ETotal,
	}
	expenses = append(expenses, total)

	// Search table for all expenses created the user
	var exps []models.Entertainment
	result = initializers.DB.Find(&exps, "username = ?", username)

	if result.RowsAffected == 0 {
		return expenses
	}

	// Convert expense struct to new struct to be used by front-end
	for _, exp := range exps {
		reformat := expense{
			ExpenseType: "Entertainment",
			ExpenseName: exp.ExpenseName,
			Amount:      exp.Amount,
		}
		expenses = append(expenses, reformat)
	}
	return expenses
}

func getHealth(username string) []interface{} {

	expenses := []interface{}{}

	// Search user table for expense total
	var user models.User
	result := initializers.DB.First(&user, "username = ?", username)

	total := total{
		ExpenseType: "Health",
		Total:       user.HTotal,
	}
	expenses = append(expenses, total)

	// Search table for all expenses created the user
	var exps []models.Health
	result = initializers.DB.Find(&exps, "username = ?", username)

	if result.RowsAffected == 0 {
		return expenses
	}

	// Convert expense struct to new struct to be used by front-end
	for _, exp := range exps {
		reformat := expense{
			ExpenseType: "Health",
			ExpenseName: exp.ExpenseName,
			Amount:      exp.Amount,
		}
		expenses = append(expenses, reformat)
	}
	return expenses
}

func getAllExpenses(username string) []interface{} {
	all := []interface{}{}

	all = append(all, getHomeUts(username), getTravel(username), getFood(username),
		getEnt(username), getHealth(username))
	return all
}