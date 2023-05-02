package db

import (
    "fmt"
    "apitemplate/config"
    "database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)


func GetCurrentDb(currentConfiguration config.Configuration) sql.DB {
    conectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
        currentConfiguration.DbUser,
        currentConfiguration.DbPassword,
        currentConfiguration.DbHost,
        currentConfiguration.DbPort,
        currentConfiguration.DbName,
    )

    db, err := sql.Open("mysql", conectionString)
    if err != nil {
        panic(err)
    }

    // See "Important settings" section.
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    return *db
}


func CloseCurrentDb(db sql.DB) {
    db.Close()
}


// enforce implement
type IEmployeeRepository interface {
    GetEmployees() ([]Employee, error)
}
