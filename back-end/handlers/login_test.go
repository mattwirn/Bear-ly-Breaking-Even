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

func TestLogin(t *testing.T) {
	initializers.StartDatabase()
	expected := "Logged In\n"

	postBody := map[string]interface{}{
		"username": "unittest",
		"password": "testunit",
	}

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
