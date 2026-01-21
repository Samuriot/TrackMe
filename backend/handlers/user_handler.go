package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuriot/track-me/db"
	"github.com/samuriot/track-me/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetUser(c *fiber.Ctx) error {
	ctx := db.GetMongoContext(c)

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := db.GetUserByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	ctx := db.GetMongoContext(c)
	users, err := db.GetAllUsers(ctx)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(users)
}

// TODO: implement these functions later
func CreateUser(c *fiber.Ctx) error {
	ctx := db.GetMongoContext(c)

	var payload struct {
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		NetWorth    float64  `json:"net_worth"`
		Accounts    []string `json:"accounts"`
		CreditScore int      `json:"credit_score"`
		Budget      []string `json:"budget"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return fiber.ErrBadRequest
	}

	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    payload.Username,
		Email:       payload.Email,
		NetWorth:    payload.NetWorth,
		Accounts:    payload.Accounts,
		CreditScore: payload.CreditScore,
		Budget:      payload.Budget,
	}

	insertedID, err := db.CreateUser(ctx, user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": insertedID.Hex(),
	})
}

// TODO: implement these functions later
func UpdateUser(c *fiber.Ctx) error {
	ctx := db.GetMongoContext(c)

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var payload struct {
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		NetWorth    float64  `json:"net_worth"`
		Accounts    []string `json:"accounts"`
		CreditScore int      `json:"credit_score"`
		Budget      []string `json:"budget"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return fiber.ErrBadRequest
	}

	update := models.User{
		Username:    payload.Username,
		Email:       payload.Email,
		NetWorth:    payload.NetWorth,
		Accounts:    payload.Accounts,
		CreditScore: payload.CreditScore,
		Budget:      payload.Budget,
	}

	user, err := db.UpdateUser(ctx, id, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	ctx := db.GetMongoContext(c)

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := db.DeleteUserByID(ctx, id); err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
