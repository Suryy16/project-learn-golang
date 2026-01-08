package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-api/backend/internal/models"
	"todo-list-api/backend/internal/transport/rest"

	"github.com/stretchr/testify/assert"
)

// Helper function to register, login and get auth cookie
func getAuthCookie(t *testing.T, router http.Handler, email, password string) *http.Cookie {
	// Register
	registerBody := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonBody, _ := json.Marshal(registerBody)
	regReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	regReq.Header.Set("Content-Type", "application/json")
	regW := httptest.NewRecorder()
	router.ServeHTTP(regW, regReq)

	// Login
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	// Get cookie
	for _, cookie := range loginW.Result().Cookies() {
		if cookie.Name == "Authorization" {
			return cookie
		}
	}
	return nil
}

// TestGetTodosWithoutAuth tests getting todos without authentication
func TestGetTodosWithoutAuth(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Request without auth
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestGetTodosEmpty tests getting todos when user has no todos
func TestGetTodosEmpty(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "emptytodos@example.com", "password123")
	assert.NotNil(t, cookie)

	// Get todos
	req, _ := http.NewRequest("GET", "/todos", nil)
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(0), response["count"])
}

// TestCreateTodoSuccess tests successful todo creation
func TestCreateTodoSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "createtodo@example.com", "password123")
	assert.NotNil(t, cookie)

	// Create todo
	todoBody := map[string]interface{}{
		"title":       "Test Todo",
		"description": "Test Description",
		"completed":   false,
	}
	jsonBody, _ := json.Marshal(todoBody)
	req, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	todo := response["todo"].(map[string]interface{})
	assert.Equal(t, "Test Todo", todo["title"])
	assert.Equal(t, "Test Description", todo["description"])
	assert.Equal(t, false, todo["completed"])
}

// TestCreateTodoWithoutAuth tests creating todo without authentication
func TestCreateTodoWithoutAuth(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// Create todo without auth
	todoBody := map[string]interface{}{
		"title":       "Test Todo",
		"description": "Test Description",
	}
	jsonBody, _ := json.Marshal(todoBody)
	req, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestCreateTodoInvalidBody tests creating todo with invalid body
func TestCreateTodoInvalidBody(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "invalidbody@example.com", "password123")
	assert.NotNil(t, cookie)

	// Create todo with invalid JSON
	req, _ := http.NewRequest("POST", "/add", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestGetSingleTodoSuccess tests getting a specific todo
func TestGetSingleTodoSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "gettodo@example.com", "password123")
	assert.NotNil(t, cookie)

	// Create a todo first
	todoBody := map[string]interface{}{
		"title":       "Get This Todo",
		"description": "Description",
		"completed":   false,
	}
	jsonBody, _ := json.Marshal(todoBody)
	createReq, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.AddCookie(cookie)
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	todo := createResponse["todo"].(map[string]interface{})
	todoID := int(todo["id"].(float64))

	// Get the todo
	req, _ := http.NewRequest("GET", fmt.Sprintf("/todo/%d", todoID), nil)
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	responseTodo := response["todo"].(map[string]interface{})
	assert.Equal(t, "Get This Todo", responseTodo["title"])
}

// TestGetSingleTodoNotFound tests getting a non-existent todo
func TestGetSingleTodoNotFound(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "notfound@example.com", "password123")
	assert.NotNil(t, cookie)

	// Get non-existent todo
	req, _ := http.NewRequest("GET", "/todo/999999", nil)
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestUpdateTodoSuccess tests successful todo update
func TestUpdateTodoSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "updatetodo@example.com", "password123")
	assert.NotNil(t, cookie)

	// Create a todo
	todoBody := map[string]interface{}{
		"title":       "Original Title",
		"description": "Original Description",
		"completed":   false,
	}
	jsonBody, _ := json.Marshal(todoBody)
	createReq, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.AddCookie(cookie)
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	todo := createResponse["todo"].(map[string]interface{})
	todoID := int(todo["id"].(float64))

	// Update the todo
	updateBody := map[string]interface{}{
		"title":       "Updated Title",
		"description": "Updated Description",
		"completed":   true,
	}
	updateJSON, _ := json.Marshal(updateBody)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/todo/%d", todoID), bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestUpdateTodoNotFound tests updating non-existent todo
func TestUpdateTodoNotFound(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "updatenotfound@example.com", "password123")
	assert.NotNil(t, cookie)

	// Update non-existent todo
	updateBody := map[string]interface{}{
		"title":       "Updated Title",
		"description": "Updated Description",
		"completed":   true,
	}
	updateJSON, _ := json.Marshal(updateBody)
	req, _ := http.NewRequest("PUT", "/todo/999999", bytes.NewBuffer(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestDeleteTodoSuccess tests successful todo deletion
func TestDeleteTodoSuccess(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "deletetodo@example.com", "password123")
	assert.NotNil(t, cookie)

	// Create a todo
	todoBody := map[string]interface{}{
		"title":       "Delete This",
		"description": "Will be deleted",
		"completed":   false,
	}
	jsonBody, _ := json.Marshal(todoBody)
	createReq, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.AddCookie(cookie)
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	todo := createResponse["todo"].(map[string]interface{})
	todoID := int(todo["id"].(float64))

	// Delete the todo
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/todo/%d", todoID), nil)
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestDeleteTodoNotFound tests deleting non-existent todo
func TestDeleteTodoNotFound(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	cookie := getAuthCookie(t, router, "deletenotfound@example.com", "password123")
	assert.NotNil(t, cookie)

	// Delete non-existent todo
	req, _ := http.NewRequest("DELETE", "/todo/999999", nil)
	req.AddCookie(cookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestUserIsolation tests that users can only access their own todos
func TestUserIsolation(t *testing.T) {
	// Setup
	models.ConnectDatabase()
	router := rest.SetupRouter()

	// User 1 creates a todo
	cookie1 := getAuthCookie(t, router, "user1@example.com", "password123")
	assert.NotNil(t, cookie1)

	todoBody := map[string]interface{}{
		"title":       "User 1 Todo",
		"description": "Private todo",
		"completed":   false,
	}
	jsonBody, _ := json.Marshal(todoBody)
	createReq, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.AddCookie(cookie1)
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	var createResponse map[string]interface{}
	json.Unmarshal(createW.Body.Bytes(), &createResponse)
	todo := createResponse["todo"].(map[string]interface{})
	todoID := int(todo["id"].(float64))

	// User 2 tries to access User 1's todo
	cookie2 := getAuthCookie(t, router, "user2@example.com", "password123")
	assert.NotNil(t, cookie2)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/todo/%d", todoID), nil)
	req.AddCookie(cookie2)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert - User 2 should not be able to access User 1's todo
	assert.Equal(t, http.StatusNotFound, w.Code)
}
