package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func GetLogName() []string {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("couldn't loud env file because: ", err)
	}
	logNames := os.Getenv("CONFIGS")
	result := strings.Split(logNames, ",")
	fmt.Println("inside args: ")
	return result
}

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
