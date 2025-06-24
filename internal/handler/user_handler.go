package handler

import (
	"fmt"
	"log"
	"strconv"

	"my-crud/internal/db"
	"my-crud/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := db.GetAllUsers(db.Database)
	if err != nil {
		log.Printf("Ошибка при получении всех пользователей: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить список пользователей"})
	}

	return c.JSON(fiber.Map{
		"users": users,
		"count": len(users),
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Имя не может быть пустым"})
	}

	if req.Age <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Возраст не может быть отрицательным"})
	}

	userID, err := db.CreateUser(db.Database, req.Name, req.Age)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось создать"})
	}

	return c.Status(201).JSON(model.UserResponse{
		ID:      userID,
		Message: fmt.Sprintf("Пользователь %s создан", req.Name),
	})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный ID (ID должен быть числом)"})
	}

	var req model.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	if req.Age <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Возраст не может быть отрицательным"})
	}

	err = db.UpdateUserAge(db.Database, userID, req.Age)
	if err != nil {
		log.Printf("Ошибка при обновлении пользователя: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить данные пользователя"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("Возраст пользователя с ID %d обновлен на %d лет", userID, req.Age)})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный ID"})
	}

	err = db.DeleteUser(db.Database, userID)
	if err != nil {
		log.Printf("Ошибка при удалении пользователя: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось удалить пользователя"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("Пользователь с ID %d удален", userID)})
}
