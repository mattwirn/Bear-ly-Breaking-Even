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
		var user models.Home_Uts
		result := initializers.DB.Where("username = ?", body.Username).Where("expensename = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

	case "Transportation":
		var user models.Trans
		result := initializers.DB.Where("username = ?", body.Username).Where("expensename = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

	case "Food":
		var user models.Food
		result := initializers.DB.Where("username = ?", body.Username).Where("expensename = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

	case "Education":
		var user models.Edu
		result := initializers.DB.Where("username = ?", body.Username).Where("expensename = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

	case "Health":
		var user models.Health
		result := initializers.DB.Where("username = ?", body.Username).Where("expensename = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Failed to delete expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Deleted\n"))
}
