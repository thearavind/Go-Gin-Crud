package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aravind741/Go-Gin-Crud/models"
)

type RegistrationResponse struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
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

	test.Run("Registering a new user to the DB", func(t *testing.T) {
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
			if _, err := ORM.Delete(&data); err != nil {
				t.Errorf("Failed to delete the new user from the DB %s", err)
			}
		}
	})
}
