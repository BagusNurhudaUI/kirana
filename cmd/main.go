package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kirana/database"
	"kirana/handler"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	zlog "github.com/rs/zerolog/log"
)

func main() {

	err := database.Connect()
	if err != nil {
		zlog.Fatal().Msgf("Database connection error: %v", err)
	}
	fmt.Println("starting the kirana server...")
	// bookCollection := db.Collection("books")
	// bookRepo := book.NewRepo(bookCollection)
	// bookService := book.NewService(bookRepo)

	app := Setup()
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	zlog.Fatal().Msgf("%v", app.Listen(":8080"))

}

type ResponseData struct {
	Message string      `json:"message"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

type Product struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "${cyan}[${time}] ${ip}  ${status} - ${red}${latency} ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		resp, err := http.Get("https://dummyjson.com/products/1")
		if err != nil {
			fmt.Println("No response from request")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body) // response body is []byte

		var product Product
		err = json.Unmarshal(body, &product)

		// Create the JSON response
		response := ResponseData{
			Message: "Welcome to kirana application server...",
			Version: "API_VERSION", // Replace with your actual API version
			Data:    product,       // Assuming you want to include the response body as a string
		}

		// Send the JSON response
		return c.Status(fiber.StatusOK).JSON(response)
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

		zlog.Info().Msg(string(jsonData))
		// Save file to root directory:
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"data": jsonData,
		})
		// return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})

	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handler.UserList)
	v1.Post("/users", handler.UserCreate)
	// Return the configured app
	return app
}
