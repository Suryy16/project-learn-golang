# Todo List API Documentation

Base URL: `http://localhost:8080`

## Table of Contents
- [Authentication](#authentication)
- [User Endpoints](#user-endpoints)
- [Todo Endpoints](#todo-endpoints)
- [Error Responses](#error-responses)

---

## Authentication

This API uses JWT tokens stored in HTTP-only cookies for authentication.

### Cookie Name
- `Authorization`: Contains JWT token

### Token Expiration
- Tokens expire after 24 hours

---

## User Endpoints

### 1. Register User

**Endpoint:** `POST /register`

**Description:** Register a new user account

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Success Response (201):**
```json
{
  "message": "User registered successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `409 Conflict`: Email already exists

---

### 2. Login

**Endpoint:** `POST /login`

**Description:** Login and receive authentication cookie

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response Headers:**
- `Set-Cookie`: Authorization token

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Invalid email or password

---

### 3. Logout

**Endpoint:** `GET /logout`

**Description:** Logout and clear authentication cookie

**Authentication:** Required

**Success Response (200):**
```json
{
  "message": "Logged out successfully"
}
```

---

### 4. Validate Token

**Endpoint:** `GET /validate`

**Description:** Validate current authentication token

**Authentication:** Required

**Success Response (200):**
```json
{
  "message": "Token is valid",
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Invalid or expired token

---

### 5. Get User

**Endpoint:** `GET /user/:id`

**Description:** Get user details by ID

**Authentication:** Required

**URL Parameters:**
- `id` (integer): User ID

**Success Response (200):**
```json
{
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Not authenticated
- `404 Not Found`: User not found

---

### 6. Update User

**Endpoint:** `PUT /user/:id`

**Description:** Update user information

**Authentication:** Required

**URL Parameters:**
- `id` (integer): User ID

**Request Body:**
```json
{
  "email": "newemail@example.com",
  "password": "newpassword123"
}
```

**Success Response (200):**
```json
{
  "message": "data berhasil diperbarui"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Not authenticated
- `404 Not Found`: User not found

---

## Todo Endpoints

### 1. Get All Todos

**Endpoint:** `GET /todos`

**Description:** Get all todos for authenticated user

**Authentication:** Required

**Success Response (200):**
```json
{
  "todos": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Milk, eggs, bread",
      "completed": false,
      "created_at": "2026-01-08T10:00:00Z",
      "updated_at": "2026-01-08T10:00:00Z",
      "user_id": 1
    },
    {
      "id": 2,
      "title": "Workout",
      "description": "30 minutes cardio",
      "completed": true,
      "created_at": "2026-01-08T11:00:00Z",
      "updated_at": "2026-01-08T11:30:00Z",
      "user_id": 1
    }
  ],
  "count": 2
}
```

**Error Responses:**
- `401 Unauthorized`: Not authenticated
- `500 Internal Server Error`: Database error

---

### 2. Get Single Todo

**Endpoint:** `GET /todo/:id`

**Description:** Get a specific todo by ID (only if it belongs to authenticated user)

**Authentication:** Required

**URL Parameters:**
- `id` (integer): Todo ID

**Success Response (200):**
```json
{
  "todo": {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, eggs, bread",
    "completed": false,
    "created_at": "2026-01-08T10:00:00Z",
    "updated_at": "2026-01-08T10:00:00Z",
    "user_id": 1
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Not authenticated
- `404 Not Found`: Todo not found or doesn't belong to user

---

### 3. Create Todo

**Endpoint:** `POST /add`

**Description:** Create a new todo

**Authentication:** Required

**Request Body:**
```json
{
  "title": "New task",
  "description": "Task description",
  "completed": false
}
```

**Success Response (201):**
```json
{
  "todo": {
    "id": 3,
    "title": "New task",
    "description": "Task description",
    "completed": false,
    "created_at": "2026-01-08T12:00:00Z",
    "updated_at": "2026-01-08T12:00:00Z",
    "user_id": 1
  }
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Not authenticated
- `500 Internal Server Error`: Database error

---

### 4. Update Todo

**Endpoint:** `PUT /todo/:id`

**Description:** Update an existing todo (only if it belongs to authenticated user)

**Authentication:** Required

**URL Parameters:**
- `id` (integer): Todo ID

**Request Body:**
```json
{
  "title": "Updated task",
  "description": "Updated description",
  "completed": true
}
```

**Success Response (200):**
```json
{
  "message": "data berhasil diperbarui"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `401 Unauthorized`: Not authenticated
- `404 Not Found`: Todo not found or doesn't belong to user

---

### 5. Delete Todo

**Endpoint:** `DELETE /todo/:id`

**Description:** Delete a todo (only if it belongs to authenticated user)

**Authentication:** Required

**URL Parameters:**
- `id` (integer): Todo ID

**Success Response (200):**
```json
{
  "message": "data berhasil dihapus"
}
```

**Error Responses:**
- `401 Unauthorized`: Not authenticated
- `404 Not Found`: Todo not found or doesn't belong to user

---

## Error Responses

### Common Error Format

All error responses follow this format:

```json
{
  "error": "Error message description",
  "message": "Additional error details"
}
```

### HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required or invalid
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource already exists
- `500 Internal Server Error`: Server error

---

## Authentication Flow

1. **Register**: `POST /register` with email and password
2. **Login**: `POST /login` with credentials → Receive JWT cookie
3. **Access Protected Routes**: Include cookie in subsequent requests
4. **Logout**: `GET /logout` to clear cookie

---

## Example Usage with cURL

### Register
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}' \
  -c cookies.txt
```

### Get All Todos
```bash
curl -X GET http://localhost:8080/todos \
  -b cookies.txt
```

### Create Todo
```bash
curl -X POST http://localhost:8080/add \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{"title":"New task","description":"Task description","completed":false}'
```

### Update Todo
```bash
curl -X PUT http://localhost:8080/todo/1 \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{"title":"Updated task","description":"Updated description","completed":true}'
```

### Delete Todo
```bash
curl -X DELETE http://localhost:8080/todo/1 \
  -b cookies.txt
```

### Logout
```bash
curl -X GET http://localhost:8080/logout \
  -b cookies.txt
```

---

## Notes

- All timestamps are in ISO 8601 format
- User IDs are automatically assigned from authenticated user
- Todos are automatically linked to the authenticated user
- Users can only access their own todos
- Passwords are hashed using bcrypt before storage
- JWT tokens expire after 24 hours
