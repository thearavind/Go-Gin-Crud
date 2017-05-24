package handlers

import (
	"net/http"
	"bytes"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"../models"
)

type RegistrationResponse struct {
	Status   int `json:"status"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	UserId   int `json:"user_id"`
}

func TestRegistrationHandler(test *testing.T) {
	test.Run("Registration with an existing user email", func(t *testing.T) {
		params := []byte(`{"email":"aravind@aravind.com", "password": "Password", "user_name": "Aravind"}`)
		request, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(params))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		Router.ServeHTTP(response, request)
		if response.Code != http.StatusConflict {
			t.Errorf("Invalid response code: %s", response.Code)
		}
	})

	/* TODO fix the bug in delete new user*/
/*	test.Run("Registering a new user to the DB", func(t *testing.T) {
		var regResponse RegistrationResponse
		params := []byte(`{"email":"aravind@india.com", "password": "Password", "user_name": "Aravindhan"}`)
		request, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(params))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		Router.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("Invalid response code: %s", response.Code)
		}
		if err := json.NewDecoder(response.Body).Decode(&regResponse); err != nil {
			t.Errorf("Failed to decode the json response %s", err)
		} else {
			data := models.Users{UserId: regResponse.UserId}
			_ = ORM.Read(&data)
			if _, err := ORM.Delete(&data); err != nil {
				t.Errorf("Failed to delete the new user from the DB %s", err)
			}
		}
	})*/
}
