package prometheus_metrics

import (
	"log"
	"logprom/internal/regex_extractor"
	"strconv"
	"strings"
)

func UpdateMetric(metrics map[string]string, registredMetrics map[string]rgx_extract.Metric) {
	for name := range metrics {
		inf := strings.Split(name, "_")
		typeValue := inf[0]
		nameValue := inf[1]
		if typeValue == "M" {
			metricType := inf[2]
			switch metricType {
			case "gauge":
				labelsValuesList := []string{}
				for labelName := range metrics {
					detail := strings.Split(labelName, "_")
					if detail[0] == "L" && detail[1] == nameValue {
						labelsValuesList = append(labelsValuesList, metrics[labelName])
					}
				}
				metricValue := metrics[name]
				parsedMetricValue, err := strconv.ParseFloat(metricValue, 64)
				if err != nil {
					log.Fatal("couldn't parse metric value because: ", err)
				}
				myMetric := registredMetrics[nameValue]
				myMetric.Gauge.WithLabelValues(labelsValuesList...).Set(parsedMetricValue)
			case "counter":
				labelsValuesList := []string{}
				for labelName := range metrics {
					detail := strings.Split(labelName, "_")
					if detail[0] == "L" && detail[1] == nameValue {
						labelsValuesList = append(labelsValuesList, metrics[labelName])
					}
				}
				metricValue := metrics[name]
				parsedMetricValue, err := strconv.ParseFloat(metricValue, 64)
				if err != nil {
					log.Fatal("couldn't parse metric value because: ", err)
				}
				myMetric := registredMetrics[nameValue]
				myMetric.Counter.WithLabelValues(labelsValuesList...).Add(parsedMetricValue)
				// can add here additional metric types cases (metrics definition is inside rgx_extractor package)
			}
		}
	}
}
