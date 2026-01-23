package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/samuriot/track-me/models"
	"github.com/samuriot/track-me/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Payload struct {
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		NetWorth    float64  `json:"net_worth"`
		Accounts    []string `json:"accounts"`
		CreditScore int      `json:"credit_score"`
		Budget      []string `json:"budget"`
}

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	ctx := c.Context()
	idHex := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	}

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "User Not Found In DB")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	users, err := h.service.GetAllUsers(ctx)

	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Users Not Found in DB")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var payload Payload
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

	err := h.service.CreateUser(ctx, &user)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

// TODO: implement these functions later
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var payload Payload

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

	user, err := h.service.UpdateUser(ctx, id, &update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	ctx := c.Context()

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.service.DeleteUserByID(ctx, id); err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
