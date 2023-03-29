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
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

	case "Trans":
		var user models.Trans
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

	case "Food":
		var user models.Food
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

	case "Edu":
		var user models.Edu
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

	case "Health":
		var user models.Health
		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.ExpenseName).Where("amount = ?", body.Amount).Delete(&user)

		if result.Error != nil {
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Failed to delete expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Deleted\n"))
}
