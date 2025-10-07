package main

import (
	"log"

	"github.com/joho/godotenv"

	"my-crud/config"
	"my-crud/internal/db"
	"my-crud/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not set")
	}
	cfg := config.LoadConfig()

	db.InitDatabase(cfg)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("Error: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
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

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API is working!", "status": "ok"})
	})

	api := app.Group("/api")
	users := api.Group("/users")

	users.Get("/", handler.GetAllUsersHandler)
	users.Post("/", handler.CreateUserHandler)
	users.Put("/:id", handler.UpdateUserHandler)
	users.Delete("/:id", handler.DeleteUserHandler)
}
