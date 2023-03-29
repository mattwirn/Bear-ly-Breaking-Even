package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
)

/*
func TestSignup(t *testing.T) {
	initializers.StartDatabase("../.env")
	expectations := [2]string{"Signed up\n", "Username already taken\n"}

	for i := 0; i < 2; i++ {
		expected := expectations[i]

		postBody := map[string]interface{}{
			"username": "unittest",
			"password": "testunit",
		}
		body, _ := json.Marshal(postBody)
		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
		w := httptest.NewRecorder()
		Signup(w, req)
		res := w.Result()

		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if string(data) != expected {
			t.Errorf("Expected %s but got %v", expected, string(data))
		}
	}
}

func TestLogin(t *testing.T) {
	// initializers.StartDatabase("../.env")
	bodies := [3]map[string]interface{}{
		{
			"username": "unittest",
			"password": "testunit",
		},
		{
			"username": "unittest2",
			"password": "testunit",
		},
		{
			"username": "unittest",
			"password": "testunit2",
		},
	}
	expectations := [3]string{"Logged In\n", "Invalid username or password\n", "Invalid username or password\n"}

	for i, b := range bodies {

		expected := expectations[i]

		postBody := b

		body, _ := json.Marshal((postBody))

		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		w := httptest.NewRecorder()

		Login(w, r)

		res := w.Result()

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != expected {
			t.Errorf("Expected %s but got %v", expected, string(data))
		}
	}
}
*/

func TestInputIncome(t *testing.T) {
	initializers.StartDatabase("../.env")
	bodies := [2]map[string]interface{}{
		{
			"username": "test",
			"amount":   6900,
		},
		{
			"username": "t3st",
			"amount":   6900,
		},
	}
	expectations := [2]string{"Income Updated\n", "Failed to find user, username does not exist\n"}

	for i, b := range bodies {

		expected := expectations[i]

		postBody := b

		body, _ := json.Marshal((postBody))

		r := httptest.NewRequest(http.MethodPost, "/dashboard/income/update", bytes.NewReader(body))
		w := httptest.NewRecorder()

		InputIncome(w, r)

		res := w.Result()

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != expected {
			t.Errorf("Expected %s but got %v", expected, string(data))
		}
	}
}

func TestAddExpense(t *testing.T) {
	//initializers.StartDatabase("../.env")
	bodies := [2]map[string]interface{}{
		{
			"Username":    "test",
			"ExpenseType": "Food",
			"ExpenseName": "chikfila",
			"Amount":      1500,
		},
		{
			"Username":    "test",
			"ExpenseType": "f0od",
			"ExpenseName": "chikfila",
			"Amount":      1500,
		},
	}
	expectations := [2]string{"Expense Added\n", "Failed to create expense, expense type not found\n"}

	for i, b := range bodies {

		expected := expectations[i]

		postBody := b

		body, _ := json.Marshal((postBody))

		r := httptest.NewRequest(http.MethodPost, "/dashboard/expense/add", bytes.NewReader(body))
		w := httptest.NewRecorder()

		AddExpense(w, r)

		res := w.Result()

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != expected {
			t.Errorf("Expected %s but got %v at index %v", expected, string(data), i)
		}
	}
}

func TestUpdateExpense(t *testing.T) {
	//initializers.StartDatabase("../.env")
	bodies := [4]map[string]interface{}{
		{
			"Username":       "test",
			"ExpenseType":    "Food",
			"OldExpenseName": "chikfila",
			"OldAmount":      1500,
			"NewExpenseName": "chik-fil-a",
			"NewAmount":      69,
		},
		{
			"Username":       "t3st",
			"ExpenseType":    "Food",
			"OldExpenseName": "chikfila",
			"OldAmount":      1500,
			"NewExpenseName": "chik-fil-a",
			"NewAmount":      69,
		},
		{
			"Username":       "test",
			"ExpenseType":    "f0od",
			"OldExpenseName": "chikfila",
			"OldAmount":      1500,
			"NewExpenseName": "chik-fil-a",
			"NewAmount":      69,
		},
		{
			"Username":       "test",
			"ExpenseType":    "Food",
			"OldExpenseName": "cheekfila",
			"OldAmount":      1500,
			"NewExpenseName": "chik-fil-a",
			"NewAmount":      69,
		},
	}
	expectations := [4]string{"Expense Updated\n", "Failed to find expense\n", "Failed to update expense, expense type not found\n", "Failed to find expense\n"}

	for i, b := range bodies {

		expected := expectations[i]

		postBody := b

		body, _ := json.Marshal((postBody))

		r := httptest.NewRequest(http.MethodPut, "/dashboard/expense/update", bytes.NewReader(body))
		w := httptest.NewRecorder()

		UpdateExpense(w, r)

		res := w.Result()

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != expected {
			t.Errorf("Expected %s but got %v", expected, string(data))
		}
	}
}

func TestDeleteExpense(t *testing.T) {
	//initializers.StartDatabase("../.env")
	bodies := [3]map[string]interface{}{
		{
			"Username":    "test",
			"ExpenseType": "Food",
			"ExpenseName": "chik-fil-a",
			"Amount":      69,
		},
		{
			"Username":    "t33st",
			"ExpenseType": "Food",
			"ExpenseName": "chikfila",
			"Amount":      1500,
		},
		{
			"Username":    "test",
			"ExpenseType": "f0od",
			"ExpenseName": "chikfila",
			"Amount":      1500,
		},
	}
	expectations := [3]string{"Expense Deleted\n", "Failed to find expense\n", "Failed to delete expense, expense type not found\n"}

	for i, b := range bodies {

		expected := expectations[i]

		postBody := b

		body, _ := json.Marshal((postBody))

		r := httptest.NewRequest(http.MethodDelete, "/dashboard/expense/delete", bytes.NewReader(body))
		w := httptest.NewRecorder()

		DeleteExpense(w, r)

		res := w.Result()

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != expected {
			t.Errorf("Expected %s but got %v with name %s", expected, string(data), postBody["Username"])
		} else {
			fmt.Printf("Successful at index %v\n", i)
		}
	}
}
