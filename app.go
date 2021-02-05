package main

import (
	"os"

	fiber "github.com/gofiber/fiber/v2"
)

// IndexResponse json pl for index page
type IndexResponse struct {
	Name string `json:"time"`
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(IndexResponse{
			Name: "Matthew",
		})
	})
	app.Listen(":"+os.Getenv("PORT"))
}
