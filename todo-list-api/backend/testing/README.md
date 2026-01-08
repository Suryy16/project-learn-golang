# Testing Documentation

## Overview
This directory contains comprehensive test suites for the Todo List API.

## Test Files

### 1. auth_test.go
Tests for authentication and user management functionality.

**Test Cases:**
- ✅ `TestRegisterSuccess` - Successful user registration
- ✅ `TestRegisterDuplicateEmail` - Registration with duplicate email
- ✅ `TestRegisterInvalidBody` - Registration with invalid request body
- ✅ `TestLoginSuccess` - Successful login with valid credentials
- ✅ `TestLoginInvalidCredentials` - Login with wrong password
- ✅ `TestLoginNonExistentUser` - Login with non-existent user
- ✅ `TestLogout` - Logout functionality
- ✅ `TestValidateToken` - Token validation with valid token
- ✅ `TestValidateTokenWithoutAuth` - Token validation without authentication

### 2. todo_test.go
Tests for todo CRUD operations and authorization.

**Test Cases:**
- ✅ `TestGetTodosWithoutAuth` - Get todos without authentication
- ✅ `TestGetTodosEmpty` - Get todos when user has no todos
- ✅ `TestCreateTodoSuccess` - Successful todo creation
- ✅ `TestCreateTodoWithoutAuth` - Create todo without authentication
- ✅ `TestCreateTodoInvalidBody` - Create todo with invalid body
- ✅ `TestGetSingleTodoSuccess` - Get specific todo by ID
- ✅ `TestGetSingleTodoNotFound` - Get non-existent todo
- ✅ `TestUpdateTodoSuccess` - Successful todo update
- ✅ `TestUpdateTodoNotFound` - Update non-existent todo
- ✅ `TestDeleteTodoSuccess` - Successful todo deletion
- ✅ `TestDeleteTodoNotFound` - Delete non-existent todo
- ✅ `TestUserIsolation` - Users can only access their own todos

## Running Tests

### Prerequisites
1. Go 1.24.6 or higher installed
2. PostgreSQL database running
3. Required dependencies installed

### Install Dependencies
```bash
go get github.com/stretchr/testify/assert
go mod tidy
```

### Run All Tests
```bash
# From the backend/testing directory
cd backend/testing
go test -v

# Or from the project root
go test ./backend/testing/... -v
```

### Run Specific Test File
```bash
# Run only authentication tests
go test -v -run TestAuth auth_test.go

# Run only todo tests
go test -v -run TestTodo todo_test.go
```

### Run Specific Test
```bash
# Run a specific test function
go test -v -run TestLoginSuccess

# Run all tests matching a pattern
go test -v -run "TestCreate.*"
```

### Run Tests with Coverage
```bash
# Generate coverage report
go test -cover ./backend/testing/...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./backend/testing/...
go tool cover -html=coverage.out
```

### Run Tests with Race Detection
```bash
go test -race ./backend/testing/...
```

## Test Scenarios

### Authentication Scenarios
1. **Registration**
   - Valid registration
   - Duplicate email registration
   - Invalid request format
   - Missing required fields

2. **Login**
   - Valid credentials
   - Invalid password
   - Non-existent user
   - Invalid request format

3. **Token Management**
   - Token generation on login
   - Token validation
   - Token expiration
   - Cookie handling

### Todo CRUD Scenarios
1. **Create**
   - Successful creation with valid data
   - Creation without authentication
   - Invalid request body
   - Auto-assignment of user_id

2. **Read**
   - Get all todos for user
   - Get specific todo by ID
   - Empty todo list
   - Access other user's todo (should fail)

3. **Update**
   - Update own todo
   - Update non-existent todo
   - Update other user's todo (should fail)
   - Invalid update data

4. **Delete**
   - Delete own todo
   - Delete non-existent todo
   - Delete other user's todo (should fail)

### Security Scenarios
1. **Authorization**
   - Protected routes require authentication
   - Invalid token rejection
   - Expired token handling
   - Missing cookie handling

2. **Data Isolation**
   - Users can only see their own todos
   - Users cannot modify others' todos
   - Users cannot delete others' todos

## Expected Test Results

### Success Criteria
- All authentication flows work correctly
- CRUD operations function as expected
- Authorization properly enforced
- Data isolation between users maintained
- Error handling returns appropriate status codes

### Status Codes
- `200 OK` - Successful GET, PUT requests
- `201 Created` - Successful POST requests
- `400 Bad Request` - Invalid request format
- `401 Unauthorized` - Authentication required/failed
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server errors

## Troubleshooting

### Database Connection Issues
If tests fail with database connection errors:
1. Ensure PostgreSQL is running
2. Check database credentials in setup.go
3. Verify database exists: `todolistdb`

### Test Failures
If specific tests fail:
1. Check if database tables exist
2. Run migrations: `go run cmd/app/main.go`
3. Clear test data between runs
4. Check for port conflicts (default: 8080)

### Clean Test Environment
To reset the test database:
```sql
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;
```
Then restart the application to run migrations.

## Best Practices

1. **Isolation**: Each test should be independent
2. **Cleanup**: Tests should not leave residual data
3. **Naming**: Use descriptive test names
4. **Coverage**: Aim for >80% code coverage
5. **Performance**: Tests should run quickly (<5s total)

## CI/CD Integration

### GitHub Actions Example
```yaml
name: Run Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: password
          POSTGRES_DB: todolistdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.24.6
      - run: go test -v ./backend/testing/...
```

## Contributing

When adding new features:
1. Write tests first (TDD approach)
2. Ensure all existing tests pass
3. Add test documentation to this README
4. Maintain >80% code coverage
