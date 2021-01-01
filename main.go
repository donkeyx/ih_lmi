package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&User{"John", 20})
		// => {"name":"John", "age":20}
	})

	log.Fatal(app.Listen(":3000"))
}
