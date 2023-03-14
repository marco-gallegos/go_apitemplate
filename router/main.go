
package router

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

const routePrefix string = "employee"

type Employee struct {
    Id          int     `json:"id"`
    Name        string  `json:"name"`
}


func getEmployees(c *fiber.Ctx) error {
    var employees []Employee  = []Employee{
        {
            Id:12,
            Name:"marco g",
        },
    }
    return c.JSON(employees)
}


func EmployeeRouter(router fiber.Router){
    router.Get(fmt.Sprintf("%s", routePrefix), getEmployees)
}


func RegisterAllInternalRoutes(router fiber.Router) (bool, error) {
    EmployeeRouter(router)
    return true, nil
}
