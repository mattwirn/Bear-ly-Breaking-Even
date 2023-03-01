package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/handlers"
)

func TestLogin(t *testing.T) {
	expected := "Logged In\n"

	postBody := map[string]interface{}{
		"username": "unittest",
		"password": "testunit",
	}

	body, _ := json.Marshal((postBody))

	r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handlers.Login(w, r)

	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
