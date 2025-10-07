package handler

import (
	"fmt"
	"log"
	"strconv"

	"my-crud/internal/db"
	"my-crud/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := db.GetAllUsers(db.Database)
	if err != nil {
		log.Printf("Failed to get all users: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user list"})
	}

	return c.JSON(fiber.Map{
		"users": users,
		"count": len(users),
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("field '%s' failed validation rule '%s'", e.Field(), e.Tag()))
		}
		return c.Status(400).JSON(fiber.Map{
			"error":  "validation failed",
			"fields": errors,
		})
	}

	userID, err := db.CreateUser(db.Database, req.Name, req.Age)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to create user"})
	}

	return c.Status(201).JSON(model.UserResponse{
		ID:      userID,
		Message: fmt.Sprintf("User %s created", req.Name),
	})
}

func UpdateUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID (ID must be a number)"})
	}

	var req model.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid age value"})
	}

	if err := db.UpdateUserAge(db.Database, userID, req.Age); err != nil {
		log.Printf("Failed to update user: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user data"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("User %d age updated to %d", userID, req.Age)})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := db.DeleteUser(db.Database, userID); err != nil {
		log.Printf("Failed to delete user: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"Message": fmt.Sprintf("User %d deleted", userID)})
}
