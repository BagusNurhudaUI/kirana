package main

import (
	"fmt"
	"kirana/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// db, cancel, err := databaseConnection()
	// if err != nil {
	// 	log.Fatal("Database Connection Error $s", err)
	// }
	// fmt.Println("Database connection success!")
	fmt.Println("starting the kirana server...")
	// bookCollection := db.Collection("books")
	// bookRepo := book.NewRepo(bookCollection)
	// bookService := book.NewService(bookRepo)

	app := Setup()
	// api := app.Group("/api")
	// routes.BookRouter(api, bookService)
	// defer cancel()
	log.Fatal(app.Listen(":8080"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Welcome to kirana application server...",
			"version": config.Config("API_VERSION"),
			// "data":    nil,
		})
	})
	// Return the configured app
	return app
}
