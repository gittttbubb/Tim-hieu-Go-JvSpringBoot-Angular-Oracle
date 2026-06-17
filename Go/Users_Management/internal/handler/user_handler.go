package handler

import (
	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/service"
	"go-rbac-system/internal/validator"
	"go-rbac-system/pkg/response"
)

type UserHandler struct {
	userService service.UserService
	overrideService service.UserPermissionOverrideService
}

func NewUserHandler(
	userService service.UserService,
	overrideService service.UserPermissionOverrideService,
) *UserHandler {
	return &UserHandler{
		userService:     userService,
		overrideService: overrideService,
	}
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.userService.List()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	return response.Success(c, user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	createdBy := c.Locals(constants.ContextUserID).(string)

	err := h.userService.Create(&req, createdBy)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "user created successfully",
	})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	err := h.userService.Update(id, &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "user updated successfully",
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userService.Delete(id)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "user deleted successfully",
	})
}

func (h *UserHandler) Lock(c *fiber.Ctx) error {
	id := c.Params("id")

	req := dto.UpdateUserRequest{}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	req = dto.UpdateUserRequest{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		RoleID:   user.RoleID,
		Status:   constants.UserStatusLocked,
	}

	err = h.userService.Update(id, &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "user locked",
	})
}

func (h *UserHandler) Unlock(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error())
	}

	req := dto.UpdateUserRequest{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		RoleID:   user.RoleID,
		Status:   constants.UserStatusActive,
	}

	err = h.userService.Update(id, &req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "user unlocked",
	})
}

func (h *UserHandler) GetOverrides(c *fiber.Ctx) error {
	userID := c.Params("id")

	data, err := h.overrideService.GetByUserID(userID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, data)
}

func (h *UserHandler) AssignOverride(c *fiber.Ctx) error {
	var req dto.UserPermissionOverrideRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	createdBy := c.Locals(constants.ContextUserID).(string)

	err := h.overrideService.Assign(&req, createdBy)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.Map{
		"message": "permission override assigned",
	})
}

func (h *UserHandler) RemoveOverride(c *fiber.Ctx) error {
    id := c.Params("id")

    err := h.overrideService.Delete(id)
    if err != nil {
        return response.Error(
            c,
            fiber.StatusBadRequest,
            err.Error(),
        )
    }

    return response.Success(c, fiber.Map{
        "message": "permission override removed",
    })
}