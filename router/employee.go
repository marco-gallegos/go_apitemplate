package router

import (
	"apitemplate/handlers"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const routePrefix string = "employee"

type EmployeeRouter struct {
    currentDb sql.DB
}


func (er *EmployeeRouter) getEmployees(c *fiber.Ctx) error {
    employees := handlers.GetEmployees(er.currentDb)
    return c.JSON(employees)
}


func RegisterEmployeeRoutes(router fiber.Router, currentDb sql.DB){
    employeeRouter := EmployeeRouter{ currentDb: currentDb }
    router.Get(fmt.Sprintf("%s", routePrefix), employeeRouter.getEmployees)
}
