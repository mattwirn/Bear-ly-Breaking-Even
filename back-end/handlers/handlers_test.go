package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
)

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
