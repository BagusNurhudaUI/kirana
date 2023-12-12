package main

import (
	"encoding/json"
	"fmt"
	"kirana/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	// api := app.Group("/api")
	// routes.BookRouter(api, bookService)
	// defer cancel()
	log.Fatal(app.Listen(":8080"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Welcome to kirana application server...",
			"version": config.Config("API_VERSION"),
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		fileInfo := struct {
			FileName    string `json:"filename"`
			Size        int64  `json:"size"`
			ContentType string `json:"content_type"`
		}{
			FileName:    file.Filename,
			Size:        file.Size,
			ContentType: file.Header.Get("Content-Type"),
		}

		// Convert the struct to JSON
		jsonData, err := json.Marshal(fileInfo)
		if err != nil {
			return err
		}

		// Print or send the JSON response
		fmt.Println(string(jsonData))
		// Save file to root directory:
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"data": jsonData,
		})
		// return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})
	// Return the configured app
	return app
}
