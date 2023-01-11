package prometheus_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"strings"
)

type metrics interface {
	InitMetric()
	createMetric()
}

var metricsNames = map[string]struct{}{}

func createMetric(metricName string, metricType string, metricValue string, metricLabels []string, labelsValue []string) {
	switch metricType {
	case "guage":
		metric := prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: metricName,
			},
			metricLabels,
		)
		metricSetValue, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			log.Fatal("couldn't parse metric value into float")
		}
		prometheus.MustRegister(metric)

		metric.WithLabelValues("", "").Set(metricSetValue)
	}
}

func InitMetric(metrics map[string]string) {
	for name := range metrics {
		inf := strings.Split(name, "_")
		switch inf[0] {
		case "M":
			metricName := inf[1]
			metricType := inf[2]
			metricValue := metrics[name]
			if _, exist := metricsNames[metricName]; !exist {
				createMetric(metricName, metricType, metricValue)
			}
		}
		metricName := metricInf[1]
	}
}
