package restful

import (
	"fmt"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) parseToken(c *fiber.Ctx) (port.AuthParams, error) {
	authorizationHeader := c.Get(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		msg := "authorization header not provided"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		msg := "invalid authorization format"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	authorizationType := fields[0]
	if authorizationType != authorizationTypeToken {
		msg := fmt.Sprintf("authorization type %s not supported", authorizationType)
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return port.AuthParams{}, err
	}

	token := fields[1]
	payload, err := s.usecase.TokenMaker().VerifyToken(token)
	if err != nil {
		return port.AuthParams{}, err
	}
	return port.AuthParams{Token: token, Payload: payload}, nil
}

func hasToken(c *fiber.Ctx) bool {
	authorizationHeader := c.Get(authorizationHeaderKey)
	return len(authorizationHeader) > 0
}

func getAuthArg(c *fiber.Ctx) (port.AuthParams, error) {
	arg := c.Locals(authorizationArgKey)
	if arg == nil {
		return port.AuthParams{}, exception.New(exception.TypePermissionDenied, "no authorization arguments provided", nil)
	}
	authArg, ok := arg.(port.AuthParams)
	if !ok {
		return port.AuthParams{}, exception.New(exception.TypePermissionDenied, "invalid authorization arguments", nil)
	}
	return authArg, nil
}
