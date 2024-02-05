package restful

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUc    port.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(uu port.UserUsecase, validator *validator.Validate) *userHandler {
	return &userHandler{userUc: uu, validator: validator}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	req := RegisterRequest{}

	if err := c.BodyParser(&req); err != nil {
		return errorHandler(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return errorHandler(c, err)
	}

	user, err := h.userUc.Register(port.RegisterParams{
		User: domain.User{
			Email:    req.Email,
			Username: req.Username,
			Password: req.Password,
		},
	})
	if err != nil {
		return errorHandler(c, err)
	}

	res := UserResponse{serializeUser(user)}
	return c.Status(fiber.StatusCreated).JSON(res)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	req := LoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return errorHandler(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return errorHandler(c, err)
	}

	user, err := h.userUc.Login(port.LoginParams{
		User: domain.User{
			Email:    req.Email,
			Password: req.Password,
		},
	})
	if err != nil {
		return errorHandler(c, err)
	}

	res := UserResponse{serializeUser(user)}
	return c.JSON(res)
}

func (h *userHandler) Current(c *fiber.Ctx) error {
	arg, err := getAuthArg(c)
	if err != nil {
		return errorHandler(c, err)
	}

	user, err := h.userUc.Current(arg)
	if err != nil {
		return errorHandler(c, err)
	}

	res := UserResponse{serializeUser(user)}
	return c.JSON(res)
}

type UpdateRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty"`
	Bio      string `json:"bio" validate:"omitempty"`
	Image    string `json:"image" validate:"omitempty"`
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	arg, err := getAuthArg(c)
	if err != nil {
		return errorHandler(c, err)
	}

	req := UpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return errorHandler(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return errorHandler(c, err)
	}

	user, err := h.userUc.Update(port.UpdateUserParams{
		AuthArg: arg,
		User: domain.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Bio:      &req.Bio,
			Image:    &req.Image,
		},
	})
	if err != nil {
		return errorHandler(c, err)
	}

	res := UserResponse{serializeUser(user)}
	return c.JSON(res)
}
