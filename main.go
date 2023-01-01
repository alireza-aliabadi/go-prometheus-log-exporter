package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	rwfiles "logprom/internal"
	"logprom/internal/env"
	logmetric "logprom/internal/logmetrics"
	"net/http"
)

var logPath = env.GetLogPath()
var responseLogPath = fmt.Sprintf("%s/responses.log", logPath)
var requestLogPath = fmt.Sprintf("%s/requests.log", logPath)
var errorsLogPath = fmt.Sprintf("%s/errors.log", logPath)

func ResponseGaugeHandler(w http.ResponseWriter, req *http.Request) {
	rwfiles.ReadFile(responseLogPath, "login")
}
func RequestGaugeHandler(w http.ResponseWriter, req *http.Request) {
	rwfiles.ReadFile(requestLogPath)
}
func ErrorGaugeHandler(w http.ResponseWriter, req *http.Request) {
	rwfiles.ReadFile(errorsLogPath, "error-count")
}

func main() {
	prometheus.MustRegister(logmetric.SuccessLogGauge)
	prometheus.MustRegister(logmetric.FailedLogGauge)
	prometheus.MustRegister(logmetric.ErrCounter)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/responses_log", ResponseGaugeHandler)
	http.HandleFunc("/requests_log", RequestGaugeHandler)
	http.HandleFunc("/errors_log", ErrorGaugeHandler)
	http.ListenAndServe(":3030", nil)
	// 	cmd.Execute()
}
