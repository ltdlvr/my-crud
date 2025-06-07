package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

//теперь мейн с увлекательными комментариями

var db *sql.DB

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateUserRequest struct {
	Age int `json:"age"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("нету.env")
	}

	InitDatabase()
	createTable()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("Ошибка: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Внутренняя ошибка сервера",
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New())

	setupRoutes(app)

	log.Println("Сервер запустили на порту 8080")
	log.Println("API доступно по адресам:")
	log.Println("   GET    /api/users     - получить всех пользователей")
	log.Println("   POST   /api/users     - создать пользователя")
	log.Println("   PUT    /api/users/:id - обновить возраст пользователя")
	log.Println("   DELETE /api/users/:id - удалить пользователя")

	log.Fatal(app.Listen(":8080"))
}

func InitDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка соединения с БД", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Ошибка пинга БД", err)
	} else {
		fmt.Println("УРААААААА С ПОДКЛЮЧЕНИЕМ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}
}

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API работает!", "status": "ok"})
	})

	api := app.Group("/api")
	users := api.Group("/users")

	users.Get("/", getAllUsersHandler)
	users.Post("/", createUserHandler)
	users.Put("/:id", updateUserHandler)
	users.Delete("/:id", deleteUserHandler)
}

func getAllUsersHandler(c *fiber.Ctx) error {
	users, err := GetAllUsers(db)
	if err != nil {
		log.Printf("Ошибка при получении всех пользователей: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить список пользователей"})
	}

	return c.JSON(fiber.Map{
		"users": users,
		"count": len(users),
	})
}

func createUserHandler(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Имя не может быть пустым"})
	}

	if req.Age <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Возраст не может быть отрицательным"})
	}

	userID, err := CreateUser(db, req.Name, req.Age)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось создать"})
	}

	return c.Status(201).JSON(UserResponse{
		ID:      userID,
		Message: fmt.Sprintf("Пользователь %s создан", req.Name),
	})
}

func updateUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный ID (ID должен быть числом)"})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if req.Age <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Возраст не может быть отрицательным"})
	}

	err = UpdateUserAge(db, userID, req.Age)
	if err != nil {
		log.Printf("Ошибка при обновлении пользователя: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить данные пользователя"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("Возраст пользователя с ID %d обновлен на %d лет", userID, req.Age)})
}

func deleteUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный ID"})
	}

	err = DeleteUser(db, userID)
	if err != nil {
		log.Printf("Ошибка при удалении пользователя: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось удалить пользователя"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("Пользователь с ID %d удален", userID)})
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы: ", err)
	}
	fmt.Println("Таблица готова")
}
// main branch
// main branch
// main branch
