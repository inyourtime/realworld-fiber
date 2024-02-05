package restful

import "github.com/gofiber/fiber/v2"

const (
	authorizationHeaderKey = "Authorization"
	authorizationTypeToken = "Bearer"
	authorizationArgKey    = "authorization_arg"
)

// AuthMiddleware is a function that serves as the middleware for authentication.
//
// It takes a boolean parameter autoDenied and returns a fiber.Handler.
func (s *Server) AuthMiddleware(autoDenied bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authArg, err := s.parseToken(c)
		if err != nil {
			if hasToken(c) {
				return errorHandler(c, err)
			}
			if autoDenied {
				return errorHandler(c, err)
			}
		}
		c.Locals(authorizationArgKey, authArg)
		return c.Next()
	}
}
