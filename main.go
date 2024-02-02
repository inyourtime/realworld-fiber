package main

import (
	"fmt"
	"log"
	"os"
	"realworld-go-fiber/adapter/repository/sql"
	"realworld-go-fiber/adapter/repository/sql/model"
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env config: %s", err)
		os.Exit(1)
	}
	fmt.Println(config.IsProduction())

	db, err := gorm.Open(postgres.Open(config.PostgresSource), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		DryRun:         false,
		TranslateError: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	pg, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{}, &model.Tag{})

	fmt.Println("connect db success")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		cond := domain.User{
			Email: "admin1@gmail.com",
			// Username: "admin2",
			// Password: "1234",
		}

		ur := sql.NewUserRepository(db)
		// userN, err := ur.Create(user1)
		userN, err := ur.FindOne(cond)

		if err != nil {
			fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
		return c.JSON(userN)
	})
	app.Listen(":" + config.ServerPort)
}
