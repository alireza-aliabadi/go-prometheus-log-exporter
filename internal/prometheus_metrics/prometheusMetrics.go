package prometheus_metrics

import (
	"log"
	"logprom/internal/regex_extractor"
	"strconv"
	"strings"
)

//type metric struct {
//	metricValue  string
//	metricType   string
//	metricLabels []string
//	labelsvalues []string
//}
//
//func (m *metric) UpdateLabels(label string, value string) {
//	for indx, metricLabel := range m.metricLabels {
//		if metricLabel == label {
//			m.labelsvalues[indx] = value
//			return
//		}
//	}
//	m.metricLabels = append(m.metricLabels, label)
//	m.labelsvalues = append(m.labelsvalues, value)
//}
//
//func (m *metric) addOrphans(labels []string, values []string) {
//	for indx, metricLabel := range m.metricLabels {
//		tempLabel := ""
//		tempValue := ""
//		for labelsIndx, label := range labels {
//			tempLabel = label
//			tempValue = values[labelsIndx]
//			if metricLabel == tempLabel {
//				m.labelsvalues[indx] = tempValue
//				continue
//			}
//		}
//		m.metricLabels = append(m.metricLabels, tempLabel)
//		m.labelsvalues = append(m.labelsvalues, tempValue)
//	}
//	fmt.Println("metric labels in orphans func --> \n", m.metricLabels)
//	fmt.Println("labels values orphans func --> \n", m.labelsvalues)
//}
//
//func pushMetricToPrometheus(metricList map[string]*metric) {
//	for metricName := range metricList {
//		switch metricList[metricName].metricType {
//		case "gauge":
//			prometheusMetric := prometheus.NewGaugeVec(
//				prometheus.GaugeOpts{
//					Name: metricName,
//				},
//				metricList[metricName].metricLabels,
//			)
//			parsingMetricvalue, err := strconv.ParseFloat(metricList[metricName].metricValue, 64)
//			if err != nil {
//				log.Fatal("couldn't parse metric value into float", err)
//			}
//			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
//			prometheus.MustRegister(prometheusMetric)
//		case "counter":
//			prometheusMetric := prometheus.NewCounterVec(
//				prometheus.CounterOpts{
//					Name: metricName,
//				},
//				metricList[metricName].metricLabels,
//			)
//			metricValue := metricList[metricName].metricValue
//			parsingMetricvalue, err := strconv.ParseFloat(metricValue, 64)
//			if metricValue != "" && err != nil {
//				log.Fatal("couldn't parse metric value into float", err)
//			} else if metricValue == "" {
//				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Inc()
//			} else {
//				prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Add(parsingMetricvalue)
//			}
//			prometheus.MustRegister(prometheusMetric)
//			// add other arbitary metric types here
//		default:
//			prometheusMetric := prometheus.NewGaugeVec(
//				prometheus.GaugeOpts{
//					Name: metricName,
//				},
//				metricList[metricName].metricLabels,
//			)
//			parsingMetricvalue, err := strconv.ParseFloat(metricList[metricName].metricValue, 64)
//			if err != nil {
//				log.Fatal("couldn't parse metric value into float", err)
//			}
//			prometheusMetric.WithLabelValues(metricList[metricName].labelsvalues...).Set(parsingMetricvalue)
//			prometheus.MustRegister(prometheusMetric)
//		}
//	}
//}
//
//var metricsList = map[string]*metric{}

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
			}
		}
	}
}
