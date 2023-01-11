package env

import (
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
	return result
}

func GetLogPath() string {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("can't load .env file because:", err)
	}

	logs_path := os.Getenv("LOGS_PATH")
	return logs_path
}

func GetRegexPattern() string {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("can't load .env file because:", err)
	}
	regexPattern := os.Getenv("REGEX")
	return regexPattern
}
