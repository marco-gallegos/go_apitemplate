package router

import (
	"apitemplate/config"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RegisterAllInternalRoutes(router fiber.Router, currentDb sql.DB, currentConfiguration config.Configuration) (bool, error) {
    RegisterEmployeeRoutes(router, currentDb)
    return true, nil
}
