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
	rwfiles.ReadFile(responseLogPath, "login")
}
func RequestGaugeHandler() {
	rwfiles.ReadFile(requestLogPath)
}
func ErrorGaugeHandler() {
	rwfiles.ReadFile(errorsLogPath, "error-count")
}

func methodCaller(arg string) {
	switch arg {
	case "responses":
		go ResponseGaugeHandler()
	case "requests":
		go RequestGaugeHandler()
	case "errors":
		go ErrorGaugeHandler()
	}
}

var rootCmd = &cobra.Command{
	Use:   "call log readers",
	Short: "call logs exporter function for automation",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			for {
				args := env.GetLogName()
				fmt.Println("args: ", args)
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
