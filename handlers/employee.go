package handlers

import (
	"apitemplate/db"
    "database/sql"
    //"fmt"
)


func GetEmployees(currentDb sql.DB) []db.Employee {
    var EmployeeRepository db.IEmployeeRepository = &db.EmployeeRepository{ CurrentDb: currentDb }

    res,_ := EmployeeRepository.GetEmployees()

    //fmt.Println(res)

    return res
}
