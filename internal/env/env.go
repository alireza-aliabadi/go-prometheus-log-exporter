package env

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// --------------------------- for using env file -------------------------

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

	logsPath := os.Getenv("LOGS_PATH")
	return logsPath
}

func GetRegexPattern() string {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal("can't load .env file because:", err)
	}
	regexPattern := os.Getenv("REGEX")
	return regexPattern
}

// ------------------------------- for using env yaml file ---------------------------

type logConf struct {
	Name  string
	File  string
	Regex string
}

type Conf struct {
	Logs  []string
	Confs []logConf
}

var c Conf

func GetEnvValues() *Conf {
	viper.SetConfigName("env.yml")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("viper couldn't read env.yml file: ", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("couldn't parse yaml confs: ", err)
	}
	return &c
}
