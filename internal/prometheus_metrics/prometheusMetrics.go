package prometheus_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
	"strings"
)

type metric struct {
	metricValue  string
	metricType   string
	metricLabels []string
	labelsvalues []string
}

func (m *metric) UpdateLabels(label string, value string) {
	m.metricLabels = append(m.metricLabels, label)
	m.labelsvalues = append(m.labelsvalues, value)
}

func (m *metric) addOrphans(labels []string, values []string) {
	m.metricLabels = append(m.metricLabels, labels...)
	m.labelsvalues = append(m.labelsvalues, values...)
}

func pushMetricToPrometheus(metricList map[string]*metric) {
	for metricName := range metricList {
		switch metricList[metricName].metricType {
		case "gauge":
			prometheusMetric := prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: metricName,
				},
				metricList[metricName].metricLabels,
			)
			parsingMetricvalue, err := strconv.ParseFloat(metricList[metricName].metricValue, 64)
			if err != nil {
				log.Fatal("couldn't parse metric value into float", err)
			}
			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
			prometheus.MustRegister(prometheusMetric)
		case "counter":
			prometheusMetric := prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: metricName,
				},
				metricList[metricName].metricLabels,
			)
			metricValue := metricList[metricName].metricValue
			parsingMetricvalue, err := strconv.ParseFloat(metricValue, 64)
			if metricValue != "" && err != nil {
				log.Fatal("couldn't parse metric value into float", err)
			} else if metricValue == "" {
				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Inc()
			} else {
				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Add(parsingMetricvalue)
			}
			prometheus.MustRegister(prometheusMetric)
			// add other arbitary metric types here
		default:
			prometheusMetric := prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: metricName,
				},
				metricList[metricName].metricLabels,
			)
			parsingMetricvalue, err := strconv.ParseFloat(metricList[metricName].metricValue, 64)
			if err != nil {
				log.Fatal("couldn't parse metric value into float", err)
			}
			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
			prometheus.MustRegister(prometheusMetric)
		}
	}
}

func InitMetric(metrics map[string]string) {
	metricsList := map[string]*metric{}
	orphanLabels := []string{}
	orphanLabelsValues := []string{}
	for name := range metrics {
		inf := strings.Split(name, "_")
		switch inf[0] {
		case "M":
			metricName := inf[1]
			typeOfMetric := inf[2]
			valueOfMetric := metrics[name]
			if _, exist := metricsList[metricName]; !exist || metricsList[metricName].metricType != typeOfMetric {
				// create metric in this section
				metricsList[metricName] = &metric{
					metricValue: valueOfMetric,
					metricType:  typeOfMetric,
				}
				if len(orphanLabels) >= 1 {
					m := metricsList[metricName]
					m.addOrphans(orphanLabels, orphanLabelsValues)
					orphanLabels = nil
					orphanLabelsValues = nil
				}
			} else {
				// update value of metric
				oldMetric := metricsList[metricName]
				oldMetric.metricValue = valueOfMetric
				if len(orphanLabels) >= 1 {
					oldMetric.addOrphans(orphanLabels, orphanLabelsValues)
					orphanLabels = nil
					orphanLabelsValues = nil
				}
			}
		case "L":
			metricName := inf[1]
			labelName := inf[2]
			labelValue := metrics[name]
			if val, ok := metricsList[metricName]; ok {
				val.UpdateLabels(labelName, labelValue)
			} else {
				orphanLabels = append(orphanLabels, labelName)
				orphanLabelsValues = append(orphanLabelsValues, labelValue)
			}
		}
	}
	pushMetricToPrometheus(metricsList)
}
