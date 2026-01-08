# 📝 Todo List API

A full-stack Todo List application built with Go (Golang) backend and vanilla HTML/CSS/JavaScript frontend.

## 🚀 Features

- ✅ User authentication (Register/Login/Logout)
- ✅ JWT-based session management with HTTP-only cookies
- ✅ CRUD operations for todos
- ✅ Toggle todo completion status
- ✅ Filter todos (All/Pending/Completed)
- ✅ Real-time statistics (Total/Completed/Pending tasks)
- ✅ Responsive UI design
- ✅ RESTful API architecture
- ✅ PostgreSQL database

## 🛠️ Tech Stack

### Backend
- **Language:** Go 1.24+
- **Framework:** Gin (HTTP web framework)
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT (golang-jwt/jwt/v5)
- **Password Hashing:** Bcrypt
- **CORS:** gin-contrib/cors

### Frontend
- **HTML5**
- **CSS3**
- **Vanilla JavaScript** (ES6+)
- **Unicons** (Icon library)

## 📁 Project Structure

```
todo-list-api/
├── backend/
│   ├── api/
│   │   ├── API_DOCUMENTATION.md      # API documentation
│   │   └── postman_collection.json   # Postman collection
│   ├── build/                        # Build artifacts
│   ├── cmd/
│   │   └── app/
│   │       └── main.go              # Application entry point
│   ├── internal/
│   │   ├── models/
│   │   │   ├── setup.go             # Database setup
│   │   │   ├── todo.go              # Todo model
│   │   │   └── user.go              # User model
│   │   └── transport/
│   │       └── rest/
│   │           ├── router.go        # API routes
│   │           ├── middleware/
│   │           │   └── requireAuth.go # Auth middleware
│   │           ├── todoController/
│   │           │   └── todoController.go
│   │           └── userController/
│   │               └── userController.go
│   ├── pkg/                         # Public packages
│   ├── testing/
│   │   ├── auth_test.go
│   │   ├── todo_test.go
│   │   └── README.md
│   ├── .env                         # Environment variables
│   └── go.mod                       # Go module file
├── frontend/
│   ├── auth.html                    # Login/Register page
│   ├── index.html                   # Main todo page
│   ├── package.json                 # NPM package file
│   └── style.css                    # Styles
├── go.mod                           # Root Go module
├── .gitignore
└── README.md
```

## 🚦 Getting Started

### Prerequisites

- **Go 1.24+**
- **PostgreSQL**
- **Git**
- **Web Browser**

### Installation & Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd todo-list-api
   ```

2. **Setup PostgreSQL Database**
   ```bash
   # Using psql
   psql -U postgres
   CREATE DATABASE todolistdb;
   ```

3. **Configure Backend Environment**
   
   Edit `backend/.env` file:
   ```env
   TOKEN=your_jwt_secret_key_here
   DB=host=localhost user=postgres password=your_password dbname=todolistdb port=5432 sslmode=disable TimeZone=Asia/Jakarta
   ```

4. **Install Go Dependencies**
   ```bash
   # From project root
   go mod download
   ```

5. **Run the Backend**
   ```bash
   cd backend
   go run cmd/app/main.go
   ```
   
   Backend will run on `http://localhost:8080`

6. **Open the Frontend**
   
   Option 1: Open directly in browser
   ```bash
   # Open frontend/auth.html in your browser
   ```
   
   Option 2: Use a local server
   ```bash
   cd frontend
   python -m http.server 3000
   # Or use VS Code Live Server extension
   ```
   
   Then open `http://localhost:3000/auth.html`

## 📡 API Endpoints

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/signup` | Register new user | No |
| POST | `/login` | Login user | No |
| POST | `/logout` | Logout user | Yes |
| GET | `/validate` | Validate JWT token | Yes |

### Todo Endpoints (Protected)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/todos` | Get all user's todos | Yes |
| POST | `/todos` | Create new todo | Yes |
| GET | `/todos/:id` | Get todo by ID | Yes |
| PUT | `/todos/:id` | Update todo | Yes |
| DELETE | `/todos/:id` | Delete todo | Yes |
| PATCH | `/todos/toggle` | Toggle completion status | Yes |

### Example Requests

**Register User**
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

**Login**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -c cookies.txt \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**Create Todo**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "title": "Buy groceries",
    "description": "Milk, eggs, bread"
  }'
```

**Get All Todos**
```bash
curl -X GET http://localhost:8080/todos \
  -b cookies.txt
```

For detailed API documentation, see [API_DOCUMENTATION.md](backend/api/API_DOCUMENTATION.md)

## 🧪 Testing

### Run Tests
```bash
cd backend
go test ./testing/... -v
```

### Manual Testing with Postman
1. Import the Postman collection: `backend/api/postman_collection.json`
2. See `backend/testing/README.md` for testing guide

### Available Tests
- `auth_test.go` - Authentication tests
- `todo_test.go` - Todo CRUD tests

## 📝 Environment Variables

Backend `.env` file format:

| Variable | Description | Example |
|----------|-------------|---------|
| `TOKEN` | JWT secret key for authentication | `your_secret_key_here` |
| `DB` | PostgreSQL connection string | `host=localhost user=postgres password=yourpass dbname=todolistdb port=5432 sslmode=disable TimeZone=Asia/Jakarta` |

## 🔒 Security Features

- JWT tokens stored in HTTP-only cookies
- Passwords hashed with bcrypt
- CORS enabled via gin-contrib/cors
- Environment variables for sensitive data
- SQL injection protection via GORM
- Input validation on both frontend and backend

## 🐛 Troubleshooting

### Backend won't start
```bash
# Verify PostgreSQL is running
psql -U postgres -l

# Check environment variables
cat backend/.env

# Test database connection
psql -U postgres -d todolistdb
```

### Database connection error
- Verify PostgreSQL service is running
- Check database name, user, and password in `.env`
- Ensure database `todolistdb` exists

### Frontend can't connect to backend
- Ensure backend is running on port 8080
- Check `API_URL` in `frontend/index.html` and `frontend/auth.html`
- Verify CORS settings in backend
- Check browser console for errors

### Port already in use
```powershell
# Find process using port 8080 (Windows)
netstat -ano | findstr :8080

# Kill process (Windows)
taskkill /PID <PID> /F
```

## 📦 Dependencies

### Backend
- `gin-gonic/gin` - HTTP web framework
- `gorm.io/gorm` - ORM library
- `gorm.io/driver/postgres` - PostgreSQL driver
- `golang-jwt/jwt/v5` - JWT implementation
- `gin-contrib/cors` - CORS middleware
- `lib/pq` - PostgreSQL driver

### Frontend
- No build dependencies (vanilla JavaScript)

## 🤝 Contributing

Contributions are welcome! Please follow these steps:


Backend `.env` file format:

| Variable | Description | Example |
|----------|-------------|---------|
| `TOKEN` | JWT secret key for authentication | `your_secret_key_here` |
| `DB` | PostgreSQL connection string | `host=localhost user=postgres password=yourpass dbname=todolistdb port=5432 sslmode=disable TimeZone=Asia/Jakarta`nts

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Unicons](https://iconscout.com/unicons)

## 📧 Contact & Support

For issues, questions, or contributions:
- Open an issue on GitHub
- Submit a pull request
- Contact: suryasatya1601@gmail.com

---

**Happy Coding! 🚀**
