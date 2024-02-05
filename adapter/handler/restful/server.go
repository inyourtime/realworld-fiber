package restful

import (
	"realworld-go-fiber/adapter/repository/sql/db"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	config util.Config
	app    *fiber.App
	logger port.Logger
	db     *db.DB
}

func NewServer(config util.Config, logger port.Logger, db *db.DB) port.Server {
	server := &Server{
		config: config,
		logger: logger,
		db:     db,
	}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	app := fiber.New()

	app.Use(recover.New(), cors.New(cors.ConfigDefault))

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "route not found"})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
	})

	// register routes
	apiRouter := app.Group("/api")
	{
		server.userRouter(apiRouter)
		server.articleRouter(apiRouter)
	}

	server.app = app
}

func (server *Server) Start() error {
	return server.app.Listen(":" + server.config.ServerPort)
}

func (server *Server) userRouter(router fiber.Router) {

}

func (server *Server) articleRouter(router fiber.Router) {

}
