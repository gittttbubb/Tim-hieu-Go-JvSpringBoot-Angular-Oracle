package handler

import (
	"go-user-management/internal/dto/user"
	"go-user-management/internal/service"

	"go-user-management/internal/utils"
	"go-user-management/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService,) *UserHandler {
	return &UserHandler{userService: userService,}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}

	claims := middleware.GetClaims(c.Locals("claims"),)
	if claims == nil {
		return utils.Unauthorized(
			c,
			"invalid token",
		)
	}

	err := h.userService.CreateUser(req, claims.UserID,)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}

	return utils.Created(
		c,
		"user created successfully",
	)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return utils.NotFound(
			c,
			err.Error(),
		)
	}

	return utils.Success(
		c,
		user,
	)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.List()
	if err != nil {
		return utils.InternalServerError(
			c,
			err.Error(),
		)
	}
	response := dto.UserListResponse{Items: users, Total: len(users),}
	return utils.Success(
		c,
		response,
	)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(
			c,
			"invalid request body",
		)
	}
	err := h.userService.UpdateUser(id, req,)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}
	return utils.Success(
		c,
		"user updated successfully",
	)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.BadRequest(
			c,
			"id is required",
		)
	}
	err := h.userService.DeleteUser(id)
	if err != nil {
		return utils.BadRequest(
			c,
			err.Error(),
		)
	}
	return utils.Success(
		c,
		"user deleted successfully",
	)
}