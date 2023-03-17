package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "apitemplate/router"
    "apitemplate/config"
)


func main() {
    config.GetConfiguration()
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    router.RegisterAllInternalRoutes(app)

    app.Listen("0.0.0.0:3000")
    fmt.Println("runing on 3000")
}
