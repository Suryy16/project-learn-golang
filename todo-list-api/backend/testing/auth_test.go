package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-api/backend/internal/models"
	"todo-list-api/backend/internal/transport/rest"

	"github.com/stretchr/testify/assert"
)

// TestRegisterSuccess tests successful user registration
func TestRegisterSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Prepare request
	body := map[string]string{
		"email":    "testuser@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestRegisterDuplicateEmail tests registration with duplicate email
func TestRegisterDuplicateEmail(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// First registration
	body := map[string]string{
		"email":    "duplicate@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Second registration with same email
	req2, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// Assert - Should fail with conflict or bad request
	assert.True(t, w2.Code == http.StatusConflict || w2.Code == http.StatusBadRequest)
}

// TestRegisterInvalidBody tests registration with invalid request body
func TestRegisterInvalidBody(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestLoginSuccess tests successful login
func TestLoginSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// First register a user
	registerBody := map[string]string{
		"email":    "logintest@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(registerBody)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	regReq.Header.Set("Content-Type", "application/json")
	regW := httptest.NewRecorder()
	router.ServeHTTP(regW, regReq)

	// Now login
	loginBody := map[string]string{
		"email":    "logintest@example.com",
		"password": "password123",
	}
	loginJSON, _ := json.Marshal(loginBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Check if cookie is set
	cookies := w.Result().Cookies()
	found := false
	for _, cookie := range cookies {
		if cookie.Name == "Authorization" {
			found = true
			assert.NotEmpty(t, cookie.Value)
			break
		}
	}
	assert.True(t, found, "Authorization cookie should be set")
}

// TestLoginInvalidCredentials tests login with wrong password
func TestLoginInvalidCredentials(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Register a user
	registerBody := map[string]string{
		"email":    "wrongpass@example.com",
		"password": "correctpassword",
	}
	jsonBody, _ := json.Marshal(registerBody)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	regReq.Header.Set("Content-Type", "application/json")
	regW := httptest.NewRecorder()
	router.ServeHTTP(regW, regReq)

	// Login with wrong password
	loginBody := map[string]string{
		"email":    "wrongpass@example.com",
		"password": "wrongpassword",
	}
	loginJSON, _ := json.Marshal(loginBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestLoginNonExistentUser tests login with non-existent user
func TestLoginNonExistentUser(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Login with non-existent user
	loginBody := map[string]string{
		"email":    "nonexistent@example.com",
		"password": "password123",
	}
	loginJSON, _ := json.Marshal(loginBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	// Execute
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestLogout tests logout functionality
func TestLogout(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Register and login first
	registerBody := map[string]string{
		"email":    "logouttest@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(registerBody)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	regReq.Header.Set("Content-Type", "application/json")
	regW := httptest.NewRecorder()
	router.ServeHTTP(regW, regReq)

	loginJSON, _ := json.Marshal(registerBody)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	// Get the cookie
	var authCookie *http.Cookie
	for _, cookie := range loginW.Result().Cookies() {
		if cookie.Name == "Authorization" {
			authCookie = cookie
			break
		}
	}

	// Logout
	logoutReq, _ := http.NewRequest("GET", "/logout", nil)
	logoutReq.AddCookie(authCookie)
	logoutW := httptest.NewRecorder()
	router.ServeHTTP(logoutW, logoutReq)

	// Assert
	assert.Equal(t, http.StatusOK, logoutW.Code)
}

// TestValidateToken tests token validation
func TestValidateToken(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Register and login
	body := map[string]string{
		"email":    "validatetest@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)

	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	regReq.Header.Set("Content-Type", "application/json")
	regW := httptest.NewRecorder()
	router.ServeHTTP(regW, regReq)

	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	// Get the cookie
	var authCookie *http.Cookie
	for _, cookie := range loginW.Result().Cookies() {
		if cookie.Name == "Authorization" {
			authCookie = cookie
			break
		}
	}

	// Validate token
	req, _ := http.NewRequest("GET", "/validate", nil)
	req.AddCookie(authCookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestValidateTokenWithoutAuth tests validation without authentication
func TestValidateTokenWithoutAuth(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Validate without cookie
	req, _ := http.NewRequest("GET", "/validate", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
