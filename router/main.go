
package router

import (
    "github.com/gofiber/fiber/v2"
)

func RegisterAllInternalRoutes(router fiber.Router) (bool, error) {
    EmployeeRouter(router)
    return true, nil
}
