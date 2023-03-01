package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/handlers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
)

func TestSignup(t *testing.T) {
	initializers.StartDatabase()

	expected := "Signed up"
	postBody := map[string]interface{}{
		"username": "unittest",
		"password": "testunit",
	}
	body, _ := json.Marshal(postBody)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handlers.Signup(w, req)
	res := w.Result()

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected Signed up but got %v", string(data))
	}
}
