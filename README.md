# Go CRUD API

A REST API for user management built with Go, Fiber, and PostgreSQL.

## 🚀 Tech Stack

- **Go** — main programming language
- **Fiber** — web framework for building REST API
- **PostgreSQL** — database
- **Docker & Docker Compose** — containerization and orchestration
- **lib/pq** — PostgreSQL driver
- **godotenv** — environment variables management
- **validator** — data validation and struct field verification

## 📋 Features 

- ✅ Create user (POST)
- ✅ Get all users (GET)
- ✅ Update user age (PUT)
- ✅ Delete user (DELETE)
- ✅ CORS and logging middleware
- ✅ Error handling

## 🛠 Installation & Setup

### Prerequisites
- Docker and Docker Compose
- Git

### Clone the repository
```bash
git clone https://github.com/ltdlvr/my-crud.git
cd my-crud
```

### Run with Docker Compose
```bash
docker-compose up --build
```

The application will be available at: `http://localhost:8080`

The database will be automatically created and configured.

## 📡 API Endpoints

| Method | Endpoint         | Description     |
|--------|------------------|-----------------|
| GET    | `/api/users`     | Get all users   |
| POST   | `/api/users`     | Create new user |
| PUT    | `/api/users/:id` | Update user age |
| DELETE | `/api/users/:id` | Delete user     |

### Example Requests

**Create user:**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 25}'
```

**Get all users:**
```bash
curl http://localhost:8080/api/users
```

**Update user age:**
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"age": 30}'
```

**Delete user:**
```bash
curl -X DELETE http://localhost:8080/api/users/1
```



## 📁 Project Structure

```
.
├── Dockerfile                # Go application Docker image
├── docker-compose.yml        # Docker Compose configuration
├── env.example               # Environment variables example
├── go.mod                    # Go modules
├── go.sum                    # Module checksums
├── cmd/
│   └── my-crud/
│       └── main.go           # Application entry point
├── config/
│   └── config.go             # Loads and validates environment variables
└── internal/
    ├── db/
    │   ├── db.go             # Database connection
    │   ├── init.go           # Database initialization
    │   ├── init_test.go      # Database tests
    │   └── migrations/
    │       └── 20250627011320_create_users_table.sql # DB migration
    ├── handler/
    │   └── user_handler.go   # HTTP handlers for user operations
    └── model/
        └── user.go           # User data model
```

## ⚙️ Environment Variables

Create a `.env` file based on `env.example`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASS=your_password
DB_NAME=your_db
SSL_MODE=disable
```

## 🔧 Development

### Run locally without Docker
1. Make sure PostgreSQL is running locally
2. Create database `cruddb`
3. Copy `env.example` to `.env` and configure database connection
4. Run the application:
```bash
go run .
```

### Add new dependencies
```bash
go mod tidy
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Create a Pull Request

## 📄 License

This project is a learning project and is free to use.

---

**Author:** Andrey Skripka  
**GitHub:** [@ltdlvr](https://github.com/ltdlvr)
