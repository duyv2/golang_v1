package main

import "github.com/gofiber/fiber/v2"
import "fmt"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  fmt.Println("server listening 8080")
  app.Listen(":8080")
}
