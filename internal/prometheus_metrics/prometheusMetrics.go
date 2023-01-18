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

func pushMetricTOPrometheus(metricList map[string]metric) {
	for metricName := range metricList {
		switch metricList[metricName].metricType {
		case "gauge":
			prometheusMetric := prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: metricName,
				},
				metricList[metricName].metricLabels,
			)
			prometheus.MustRegister(prometheusMetric)
			parsingMetricvalue, err := strconv.ParseFloat(metricList[metricName].metricValue, 64)
			if err != nil {
				log.Fatal("couldn't parse metric value into float", err)
			}
			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
		case "counter":
			prometheusMetric := prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: metricName,
				},
				metricList[metricName].metricLabels,
			)
			prometheus.MustRegister(prometheusMetric)
			metricValue := metricList[metricName].metricValue
			parsingMetricvalue, err := strconv.ParseFloat(metricValue, 64)
			if metricValue != "" && err != nil {
				log.Fatal("couldn't parse metric value into float", err)
			} else if metricValue == "" {
				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Inc()
			} else {
				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Add(parsingMetricvalue)
			}
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
			prometheus.MustRegister(prometheusMetric)
			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
		}
	}
}

func InitMetric(metrics map[string]string) {
	fmt.Println("init metric recieved metrics --> ", metrics)
	metricsList := map[string]metric{}
	for name := range metrics {
		fmt.Println("inside metrics range --> ", name)
		inf := strings.Split(name, "_")
		fmt.Println("metric seperated inf --> ", inf, inf[0])
		switch inf[0] {
		case "M":
			metricName := inf[1]
			typeOfMetric := inf[2]
			valueOfMetric := metrics[name]
			if _, exist := metricsList[metricName]; !exist || metricsList[metricName].metricType != typeOfMetric {
				// create metric in this section
				metricsList[metricName] = metric{
					metricValue: valueOfMetric,
					metricType:  typeOfMetric,
				}
			} else {
				// update value of metric
				oldMetric := metricsList[metricName]
				oldMetric.metricValue = valueOfMetric
			}
		case "L":
			metricName := inf[1]
			labelName := inf[2]
			labelValue := metrics[name]
			labelsList := metricsList[metricName].metricLabels
			labelsValue := metricsList[metricName].labelsvalues
			labelsList = append(labelsList, labelName)
			labelsValue = append(labelsValue, labelValue)
		}
	}
	fmt.Println("metrics list --> ", metricsList)
	pushMetricTOPrometheus(metricsList)
}
