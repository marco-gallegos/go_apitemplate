package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "apitemplate/router"
)


func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        fmt.Println("hello marco")
        return c.SendString("Hello, World!")
    })

    router.RegisterAllInternalRoutes(app)

    app.Listen("0.0.0.0:3000")
    fmt.Println("runing on 3000")
}
