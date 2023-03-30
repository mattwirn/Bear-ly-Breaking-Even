package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
)

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username       string
		ExpenseType    string
		OldExpenseName string
		OldAmount      uint
		NewExpenseName string
		NewAmount      uint
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch body.ExpenseType {
	case "HomeUtils":
		var user models.Home_Uts

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		user.ExpenseName = body.NewExpenseName
		user.Amount = body.NewAmount
		initializers.DB.Save(&user)

	case "Trans":
		var user models.Trans

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		user.ExpenseName = body.NewExpenseName
		user.Amount = body.NewAmount
		initializers.DB.Save(&user)

	case "Food":
		var user models.Food

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		user.ExpenseName = body.NewExpenseName
		user.Amount = body.NewAmount
		initializers.DB.Save(&user)

	case "Edu":
		var user models.Edu

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		user.ExpenseName = body.NewExpenseName
		user.Amount = body.NewAmount
		initializers.DB.Save(&user)

	case "Health":
		var user models.Health

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		user.ExpenseName = body.NewExpenseName
		user.Amount = body.NewAmount
		initializers.DB.Save(&user)

	default:
		http.Error(w, "Failed to update expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Updated\n"))
}
