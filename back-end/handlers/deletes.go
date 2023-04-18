package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
)

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
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
		var expense models.Home_Uts
		var user models.User
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).First(&expense)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.HUTotal = user.HUTotal - expense.Amount
		initializers.DB.Save(&user)
		initializers.DB.Delete(&expense)

	case "Travel":
		var expense models.Travel
		var user models.User
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).First(&expense)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.TTotal = user.TTotal - expense.Amount
		initializers.DB.Save(&user)
		initializers.DB.Delete(&expense)

	case "Food":
		var expense models.Food
		var user models.User
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).First(&expense)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.FTotal = user.FTotal - expense.Amount
		initializers.DB.Save(&user)
		initializers.DB.Delete(&expense)

	case "Ent":
		var expense models.Entertainment
		var user models.User
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).First(&expense)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.ETotal = user.ETotal - expense.Amount
		initializers.DB.Save(&user)
		initializers.DB.Delete(&expense)

	case "Health":
		var expense models.Health
		var user models.User
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).First(&expense)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.HTotal = user.HTotal - expense.Amount
		initializers.DB.Save(&user)
		initializers.DB.Delete(&expense)

	default:
		http.Error(w, "Failed to delete expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Deleted\n"))
}
