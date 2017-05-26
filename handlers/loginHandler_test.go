package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aravind741/Go-Gin-Crud/models"
	"github.com/aravind741/Go-Gin-Crud/router/route_provider"
)

type LoginResponse struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

var AuthToken string
var Router = provider.ProvideRouter()
var orM = models.GetOrmObject()

func TestLoginHandler(test *testing.T) {
	test.Run("Login Test with correct password", func(t *testing.T) {
		loginResponse := LoginResponse{}
		params := []byte(`{"email":"aravind@aravind.com", "password": "Password"}`)
		request, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(params))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		Router.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Errorf("Invalid response code: %d", response.Code)
		}
		json.NewDecoder(response.Body).Decode(&loginResponse)
		AuthToken = loginResponse.Token
	})

	test.Run("Login Test with wrong password", func(t *testing.T) {
		params := []byte(`{"email":"aravind@aravind.com", "password": "password"}`)
		request, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(params))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		Router.ServeHTTP(response, request)
		if response.Code != http.StatusUnauthorized {
			t.Errorf("Invalid response code: %d", response.Code)
		}
	})

	test.Run("Login Test with not registered email", func(t *testing.T) {
		params := []byte(`{"email":"aravind@newdomain.com", "password": "Password"}`)
		request, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(params))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		Router.ServeHTTP(response, request)
		if response.Code != http.StatusUnauthorized {
			t.Errorf("Invalid response code: %d", response.Code)
		}
	})
}
