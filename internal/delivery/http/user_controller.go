package controller

import (
	e "authentication-service/internal/exception"
	"authentication-service/internal/model"
	"authentication-service/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (h *UserController) Register(c *fiber.Ctx) error {
	var req model.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return e.BadRequest(err)
	}

	user, err := h.userService.Register(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserController) GetUserByID(c *fiber.Ctx) error {
	// Retrieve user ID from the request parameters
	userID := c.Params("id")

	userId, err := strconv.Atoi(userID)
	if err != nil {
		return e.Internal(err)
	}

	// Fetch user details from the server
	response, err := h.userService.GetUserByID(userId)
	if err != nil {
		return err
	}

	// Return user details in the response
	return c.JSON(response)
}

func (h *UserController) Login(c *fiber.Ctx) error {
	var req model.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return e.BadRequest(err)
	}

	user, err := h.userService.Login(req)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserController) RefreshToken(c *fiber.Ctx) error {
	var req model.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return e.BadRequest(err)
	}

	user, err := h.userService.RefreshToken(req)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
