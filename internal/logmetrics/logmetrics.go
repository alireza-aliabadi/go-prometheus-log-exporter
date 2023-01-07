package logmetric

import (
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
	if logStatus == "200" {
		successLog.WithLabelValues(logPath).Inc()
	} else {
		failedLog.WithLabelValues(logPath).Inc()
	}
}

// add other metrics functionsl

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
