package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "apitemplate/router"
    "apitemplate/db"
    "apitemplate/config"
)

func main() {
    var currentConfiguration config.Configuration = config.GetConfiguration() 


    currentDb := db.GetCurrentDb(currentConfiguration)

    defer db.CloseCurrentDb(currentDb) // avoid opened connections

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    router.RegisterAllInternalRoutes(app, currentDb, currentConfiguration)

    app.Listen("0.0.0.0:3000")
    fmt.Println("runing on 3000")
}
