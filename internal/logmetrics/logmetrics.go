package logmetric

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
)

var SuccessLogGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "success_login_gauge",
		Help: "number of successful login attempts",
	},
	[]string{"path", "status", "resp_time"},
)

var FailedLogGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "failed_login_gauge",
		Help: "number of unsuccessful login attempts",
	},
	[]string{"path", "status", "resp_time"},
)

func LogGaugeVec(logInf map[string]string) {
	successLog := SuccessLogGauge
	failedLog := FailedLogGauge
	logStatus := logInf["status"]
	logPath := logInf["path"]
	logRespTime := logInf["response_time"]
	if logStatus == "200" {
		successLog.WithLabelValues(logPath, logStatus, logRespTime).Inc()
	} else {
		failedLog.WithLabelValues(logPath, logStatus, logRespTime).Inc()
	}
}

// add other metrics functionsl

var ResptimeGaugeVec = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "response_time_gauge",
		Help: "show response time per response",
	},
	[]string{"path"},
)

func ResponseTimeGauge(logInf map[string]string) {
	logPath := logInf["path"]
	logResponseTime, err := strconv.ParseFloat(logInf["response_time"], 64)
	if err != nil {
		log.Fatal("couldn't parse time string into float64")
	}
	respTimeGauge := ResptimeGaugeVec
	respTimeGauge.WithLabelValues(logPath).Set(logResponseTime)
}

var ErrCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "errors_occurrence_counter",
		Help: "count number of errors",
	},
	[]string{"path"},
)

func ErrCounterVec(logInf map[string]string) {
	path := logInf["path"]
	ErrCounter.WithLabelValues(path).Inc()
}
