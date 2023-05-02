package db

import (
    "database/sql"
)

// this needs to fullfill the iEmployeeRepository
type EmployeeRepository struct {
    CurrentDb sql.DB
}

func (er *EmployeeRepository) GetEmployees() ([]Employee, error) {
    employees := []Employee{}
    results, err := er.CurrentDb.Query("SELECT emp_no FROM employees limit 10")
    if err !=nil {
        panic(err.Error())
    }
    for results.Next() {
        var testtable Employee
        err = results.Scan(&testtable.Emp_no)
        if err !=nil {
            panic(err.Error())
        }
        //fmt.Println(testtable.Emp_no)
        employees = append(employees, testtable)
    }
    return employees, nil
}
