package config

import (
	"fmt"
    "github.com/joho/godotenv"
    "os"
    "strconv"
)

type Configuration struct {
    dbName string
    dbHost string
    dbPort int
    dbPassword string
    dbUser string
}

func GetConfiguration() Configuration {
    fmt.Println("runing configuration")
    godotenv.Load()

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv(("DB_PASSWORD"))
    dbName := os.Getenv(("DB_NAME"))
    dbHost := os.Getenv(("DB_HOST"))
    dbPort, err := strconv.Atoi(os.Getenv(("DB_PORT")))

    if err != nil {
        fmt.Println("port load failed using default")
    }


    //fmt.Println(fmt.Sprintf("user %s", dbUser))

    var currentConfig = Configuration{
        dbUser: dbUser,
        dbPassword: dbPassword,
        dbName: dbName,
        dbPort: dbPort,
        dbHost: dbHost,
    }

    return currentConfig
}
