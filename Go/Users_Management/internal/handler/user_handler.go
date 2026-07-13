package handler

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"

	"go-rbac-system/internal/constants"
	"go-rbac-system/internal/dto"
	"go-rbac-system/internal/model"
	"go-rbac-system/internal/service"
	"go-rbac-system/internal/validator"
	"go-rbac-system/pkg/response"
)

type UserHandler struct {
	userService service.UserService
	overrideService service.UserPermissionOverrideService
}

func NewUserHandler(userService service.UserService, overrideService service.UserPermissionOverrideService,) *UserHandler {
	return &UserHandler{
		userService: userService,
		overrideService: overrideService,
	}
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	keyword := strings.TrimSpace(c.Query("keyword"))
	users, total, err := h.userService.List(keyword,page,pageSize,)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error(),)
	}
	return response.Success(c,model.Pagination{
			Items:    users,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	)
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
		return response.ValidationError(c, err)
	}
	actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	tempPassword, err := h.userService.Create(
		&req,
		actorID,
		actor,
	)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.Map{
		"message":           "user created successfully",
		"temporaryPassword": tempPassword,
	})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	if err := validator.Validate.Struct(&req); err != nil {
		return response.ValidationError(c, err)
	}

	actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	err := h.userService.Update(
		id,
		&req,
		actorID,
		actor,
	)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.Map{
		"message": "user updated successfully",
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	err := h.userService.Delete(
		id,
		actorID,
		actor,
	)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.Map{
		"message": "user deleted successfully",
	})
}

func (h *UserHandler) Lock(c *fiber.Ctx) error {
	id := c.Params("id")
	actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	err := h.userService.LockUser(
		id,
		actorID,
		actor,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.Error(
				c,
				fiber.StatusNotFound,
				"user not found",
			)
		}
		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
	}
	return response.Success(c, fiber.Map{
		"message": "user locked",
	})
}

func (h *UserHandler) Unlock(c *fiber.Ctx) error {
	id := c.Params("id")
	actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	err := h.userService.UnlockUser(
		id,
		actorID,
		actor,
	)
		if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.Error(
				c,
				fiber.StatusNotFound,
				"user not found",
			)
		}
		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
		)
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
		return response.ValidationError(c, err)
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

func (h *UserHandler) UpdateRole(c *fiber.Ctx) error {
    userID := c.Params("id")
    var req dto.UpdateUserRoleRequest
    if err := c.BodyParser(&req); err != nil {
        return response.Error(c, fiber.StatusBadRequest, err.Error())
    }
    if err := validator.Validate.Struct(&req); err != nil {
        return response.ValidationError(c, err)
    }
    actorID := c.Locals(constants.ContextUserID).(string)
	actor := c.Locals(constants.ContextUsername).(string)

	err := h.userService.UpdateRole(
		userID,
		req.RoleID,
		actorID,
		actor,
	)
    if err != nil {
        return response.Error(c, fiber.StatusBadRequest, err.Error())
    }
    return response.Success(c, fiber.Map{
        "message": "user role updated successfully",
    })
}