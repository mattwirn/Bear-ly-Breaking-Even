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
		var expense models.Home_Uts
		var user models.User

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&expense)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}

		user.HUTotal = user.HUTotal - expense.Amount + body.NewAmount
		initializers.DB.Save(&user)
		expense.ExpenseName = body.NewExpenseName
		expense.Amount = body.NewAmount
		initializers.DB.Save(&expense)

	case "Travel":
		var expense models.Travel
		var user models.User

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&expense)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		user.TTotal = user.TTotal - expense.Amount + body.NewAmount
		initializers.DB.Save(&user)
		expense.ExpenseName = body.NewExpenseName
		expense.Amount = body.NewAmount
		initializers.DB.Save(&expense)

	case "Food":
		var expense models.Food
		var user models.User

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&expense)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}

		user.FTotal = user.FTotal - expense.Amount + body.NewAmount
		initializers.DB.Save(&user)
		expense.ExpenseName = body.NewExpenseName
		expense.Amount = body.NewAmount
		initializers.DB.Save(&expense)

	case "Ent":
		var expense models.Entertainment
		var user models.User

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&expense)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}

		user.ETotal = user.ETotal - expense.Amount + body.NewAmount
		initializers.DB.Save(&user)
		expense.ExpenseName = body.NewExpenseName
		expense.Amount = body.NewAmount
		initializers.DB.Save(&expense)

	case "Health":
		var expense models.Health
		var user models.User

		result := initializers.DB.Where("username = ?", body.Username).Where("expense_name = ?", body.OldExpenseName).Where("amount = ?", body.OldAmount).First(&expense)
		if result.Error != nil {
			http.Error(w, "Failed to find expense", http.StatusInternalServerError)
			return
		}

		result = initializers.DB.Where("username = ?", expense.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
			return
		}

		user.HTotal = user.HTotal - expense.Amount + body.NewAmount
		initializers.DB.Save(&user)
		expense.ExpenseName = body.NewExpenseName
		expense.Amount = body.NewAmount
		initializers.DB.Save(&expense)

	default:
		http.Error(w, "Failed to update expense, expense type not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expense Updated\n"))
}
