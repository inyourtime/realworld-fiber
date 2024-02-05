package restful

import (
	"net/http"
	"realworld-go-fiber/core/util/exception"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type validationErrResponse struct {
	Field string `json:"field"`
	Tag   string `json:"error"`
	Value string `json:"value,omitempty"`
}

func errorHandler(c *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case *fiber.Error:
		return c.Status(err.Code).JSON(err)
	case validator.ValidationErrors:
		var errs []validationErrResponse
		for _, e := range err {
			errs = append(errs, validationErrResponse{Field: e.Field(), Tag: e.Tag(), Value: e.Param()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":        http.StatusBadRequest,
			"message":     fiber.ErrBadRequest.Message,
			"description": errs,
		})
	case *exception.Exception:
		if !err.HasError() {
			err.AddError("exception", err.Message)
		}
		var statusCode int
		switch err.Type {
		case exception.TypeNotFound:
			statusCode = http.StatusNotFound
		case exception.TypeTokenExpired, exception.TypeTokenInvalid, exception.TypePermissionDenied:
			statusCode = http.StatusUnauthorized
		case exception.TypeValidation:
			statusCode = http.StatusUnprocessableEntity
		default:
			statusCode = http.StatusInternalServerError
		}
		return c.Status(statusCode).JSON(err)
	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}
