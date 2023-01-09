package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	rwfiles "logprom/internal"
	"logprom/internal/env"
	"net/http"
	"time"
)

var logPath = env.GetLogPath()
var responseLogPath = fmt.Sprintf("%s/responses.log", logPath)
var requestLogPath = fmt.Sprintf("%s/requests.log", logPath)
var errorsLogPath = fmt.Sprintf("%s/errors.log", logPath)

func ResponseGaugeHandler() {
	rwfiles.ReadFile(responseLogPath, "log")
}
func RequestGaugeHandler() {
	rwfiles.ReadFile(requestLogPath)
}
func ErrorGaugeHandler() {
	rwfiles.ReadFile(errorsLogPath, "error-count")
}

var methodFlag = map[string]bool{}

func methodCaller(arg string) {
	switch arg {
	case "responses":
		if ok, exists := methodFlag["responses"]; !ok || !exists {
			go ResponseGaugeHandler()
			methodFlag["responses"] = true
		}
	case "requests":
		if ok, exists := methodFlag["requests"]; !ok || !exists {
			go RequestGaugeHandler()
			methodFlag["requests"] = true
		}
	case "errors":
		if ok, exists := methodFlag["errors"]; !ok || !exists {
			go ErrorGaugeHandler()
			methodFlag["errors"] = true
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "call log readers",
	Short: "call logs exporter function for automation",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			for {
				args := env.GetLogName()
				for _, val := range args {
					methodCaller(val)
				}
				time.Sleep(2 * time.Second)
			}
		}()
		log.Fatal(http.ListenAndServe(":3030", nil))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("couldn't start service because: ", err)
	}
}
