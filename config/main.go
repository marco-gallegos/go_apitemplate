package config

import (
	"fmt"
    "github.com/joho/godotenv"
    "os"
    "strconv"
)

type Configuration struct {
    DbName string
    DbHost string
    DbPort int
    DbPassword string
    DbUser string
}

func GetConfiguration() Configuration {
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
        DbUser: dbUser,
        DbPassword: dbPassword,
        DbName: dbName,
        DbPort: dbPort,
        DbHost: dbHost,
    }

    return currentConfig
}
