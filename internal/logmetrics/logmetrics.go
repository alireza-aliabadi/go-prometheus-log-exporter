package logmetric

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
)

var LogGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "log_gauge",
		Help: "number of responses",
	},
	[]string{"path", "status"},
)

func LogGaugeVec(logInf map[string]string) {
	Log := LogGauge
	logStatus := logInf["status"]
	logPath := logInf["path"]
	Log.WithLabelValues(logPath, logStatus).Inc()
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
		return
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
