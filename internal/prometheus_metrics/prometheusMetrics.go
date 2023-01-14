package prometheus_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strings"
)

type metrics interface {
	InitMetric()
	createMetric()
}

var metricsList = map[string]map[string]string{}

func createMetric(metricList map[string]map[string]string, metricLabelsList []string, labelsValues []string) {
	for metricName := range metricList {
		metric, err := prometheus.GaugeVec.GetMetricWithLabelValues(labelsValues...)
		if err != nil {
			log.Fatal("failed to get metric --> ", err)
		}
		if metric == nil {
			switch metricList[metricName]["type"] {
			case "gauge":
				prometheus.NewGaugeVec(
					prometheus.GaugeOpts{
						Name: metricName,
					},
					metricLabelsList,
				)
			case "counter":
				prometheus.NewCounterVec(
					prometheus.CounterOpts{
						Name: metricName,
					},
					metricLabelsList,
				)
				// add other arbitary metric types here
			}
		}
	}
}

func InitMetric(metrics map[string]string) {
	metricLabels := []string{}
	labelsValue := []string{}
	for name := range metrics {
		inf := strings.Split(name, "_")
		switch inf[0] {
		case "M":
			metricName := inf[1]
			metricType := inf[2]
			metricValue := metrics[name]
			if _, exist := metricsList[metricName]; !exist {
				metricsList[metricName] = map[string]string{
					"value": metricValue,
					"type":  metricType,
				}
			} else {
				metricsList[metricName]["value"] = metricValue
			}
		case "L":
			metricName := inf[1]
			labelName := inf[2]
			labelValue := metrics[name]
			metricLabels = append(metricLabels, labelName)
			labelsValue = append(labelsValue, labelValue)
			metricsList[metricName][labelName] = labelValue
		}
	}
	createMetric(metricsList, metricLabels, labelsValue)
}
