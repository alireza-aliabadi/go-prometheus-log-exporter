package logmetric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var SuccessLogGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "success_login_gauge",
		Help: "number of successful login attempts",
	},
	[]string{"path"},
)

var FailedLogGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "failed_login_gauge",
		Help: "number of unsuccessful login attempts",
	},
	[]string{"path"},
)

func LogGaugeVec(logInf map[string]string) {
	successLog := SuccessLogGauge
	failedLog := FailedLogGauge
	logStatus := logInf["status"]
	logPath := logInf["path"]
	fmt.Println("log status --> ", logStatus)
	fmt.Println("log path --> ", logPath)
	if logStatus == "200" {
		fmt.Println("success login --->")
		successLog.WithLabelValues(logPath).Inc()
	} else {
		fmt.Println("failed login  ---->")
		failedLog.WithLabelValues(logPath).Inc()
	}
}

// add other metrics functions

var ErrCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "errors_occurrence_counter",
		Help: "count number of errors",
	},
)

func ErrCounterVec() {
	ErrCounter.Inc()
}
