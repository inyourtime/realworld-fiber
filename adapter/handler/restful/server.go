package restful

import (
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	config    util.Config
	app       *fiber.App
	logger    port.Logger
	validator *validator.Validate
	usecase   port.Usecase
}

func NewServer(config util.Config, logger port.Logger, uc port.Usecase) port.Server {
	server := &Server{
		config:  config,
		logger:  logger,
		usecase: uc,
	}

	server.setupValidator()
	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	app := fiber.New()

	app.Use(recover.New(), cors.New(cors.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
	})

	// register routes
	apiRouter := app.Group("/api")
	{
		server.userRouter(apiRouter)
		server.articleRouter(apiRouter)
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "route not found"})
	})

	server.app = app
}

func (server *Server) setupValidator() {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	server.validator = validator
}

func (server *Server) Start() error {
	return server.app.Listen(":" + server.config.ServerPort)
}

func (server *Server) userRouter(router fiber.Router) {
	userHandler := NewUserHandler(server.usecase.User(), server.validator)

	router.Post("/users", userHandler.Register)
	router.Post("/users/login", userHandler.Login)

	userRouter := router.Group("/user")
	userRouter.Use(server.AuthMiddleware(true))
	userRouter.Get("/", userHandler.Current)
	userRouter.Put("/", userHandler.Update)

	profileRouter := router.Group("/profiles")
	profileRouter.Use(server.AuthMiddleware(false))
	profileRouter.Get("/:username", userHandler.Profile)
	profileRouter.Use(server.AuthMiddleware(true))
	profileRouter.Post("/:username/follow", userHandler.Follow)
	profileRouter.Delete("/:username/follow", userHandler.UnFollow)
}

func (server *Server) articleRouter(router fiber.Router) {

}
