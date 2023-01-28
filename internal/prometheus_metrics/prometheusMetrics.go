package prometheus_metrics

import (
	"fmt"
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
	for indx, metricLabel := range m.metricLabels {
		if metricLabel == label {
			m.labelsvalues[indx] = value
			return
		}
	}
	m.metricLabels = append(m.metricLabels, label)
	m.labelsvalues = append(m.labelsvalues, value)
}

func (m *metric) addOrphans(labels []string, values []string) {
	for indx, metricLabel := range m.metricLabels {
		tempLabel := ""
		tempValue := ""
		for labelsIndx, label := range labels {
			tempLabel = label
			tempValue = values[labelsIndx]
			if metricLabel == tempLabel {
				m.labelsvalues[indx] = tempValue
				continue
			}
		}
		m.metricLabels = append(m.metricLabels, tempLabel)
		m.labelsvalues = append(m.labelsvalues, tempValue)
	}
	fmt.Println("metric labels in orphans func ---> \n", m.metricLabels)
	fmt.Println("labels values orphans func --> \n", m.labelsvalues)
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

var metricsList = map[string]*metric{}

func InitMetric(metrics map[string]string) {
	orphanLabels := []string{}
	orphanLabelsValues := []string{}
	for name := range metrics {
		inf := strings.Split(name, "_")
		switch inf[0] {
		case "M":
			metricName := inf[1]
			typeOfMetric := inf[2]
			valueOfMetric := metrics[name]
			fmt.Println("check prometheus metric existence --> ", metricsList)
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
