package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchUser(test *testing.T) {
	test.Run("Fetch current user data test", func(t *testing.T) {
		if AuthToken != "" {
			request, _ := http.NewRequest("GET", "/api/user", nil)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Authorization", AuthToken)
			response := httptest.NewRecorder()
			Router.ServeHTTP(response, request)
			if response.Code != http.StatusOK {
				t.Errorf("Invalid response code: %d", response.Code)
			}
		} else {
			t.Errorf("No Auth token")
		}
	})

}
