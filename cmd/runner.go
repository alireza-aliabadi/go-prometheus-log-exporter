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

func ResponseGaugeHandler(path string, regexPattern string) {
	rwfiles.ReadFile(path, "log", regexPattern)
}
func RequestGaugeHandler(path string, regexPattern string) {
	rwfiles.ReadFile(path, "", regexPattern)
}
func ErrorGaugeHandler(path string, regexPattern string) {
	rwfiles.ReadFile(path, "error-count", regexPattern)
}

var methodFlag = map[string]bool{}

func methodCaller(arg string, path string, regexPattern string) {
	switch arg {
	case "response":
		if ok, exists := methodFlag["response"]; !ok || !exists {
			go ResponseGaugeHandler(path, regexPattern)
			methodFlag["response"] = true
		}
	case "request":
		if ok, exists := methodFlag["request"]; !ok || !exists {
			go RequestGaugeHandler(path, regexPattern)
			methodFlag["request"] = true
		}
	case "error":
		if ok, exists := methodFlag["error"]; !ok || !exists {
			go ErrorGaugeHandler(path, regexPattern)
			methodFlag["errors"] = true
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "call log readers",
	Short: "call logs exporter function for automation",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			//for {
			//	args := env.GetLogName()
			//	for _, val := range args {
			//		methodCaller(val, path, regex)
			//	}
			//	time.Sleep(2 * time.Second)
			//}
			envInf := env.GetEnvValues()
			fmt.Println(envInf.Logs, "\n", envInf.Confs)
			for _, requiredLog := range envInf.Logs {
				for _, logConfig := range envInf.Confs {
					if requiredLog == logConfig.Name {
						methodCaller(requiredLog, logConfig.File, logConfig.Regex)
						continue
					}
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
