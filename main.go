package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"logprom/cmd"
	"net/http"
)

//var logPath = env.GetLogPath()
//var responseLogPath = fmt.Sprintf("%s/responses.log", logPath)
//var requestLogPath = fmt.Sprintf("%s/requests.log", logPath)
//var errorsLogPath = fmt.Sprintf("%s/errors.log", logPath)

//func ResponseGaugeHandler() {
//	rwfiles.ReadFile(responseLogPath, "login")
//}
//func RequestGaugeHandler() {
//	rwfiles.ReadFile(requestLogPath)
//}
//func ErrorGaugeHandler() {
//	rwfiles.ReadFile(errorsLogPath, "error-count")
//}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	//http.HandleFunc("/responses_log", ResponseGaugeHandler)
	//http.HandleFunc("/requests_log", RequestGaugeHandler)
	//http.HandleFunc("/errors_log", ErrorGaugeHandler)
	//log.Fatal(http.ListenAndServe(":3030", nil))
	cmd.Execute()
}
