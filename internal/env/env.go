package env

import (
    "os"
    "github.com/joho/godotenv"
    "log"
    "fmt"
)


func GetLogPath() string {
    fmt.Printf("inside env package")
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("can't load .env file because:", err)
    }

    logs_path := os.Getenv("LOGS_PATH")
    fmt.Println("log path in env package", logs_path)
    return logs_path
}