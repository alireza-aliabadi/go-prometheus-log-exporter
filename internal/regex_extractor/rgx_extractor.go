package rgx_extract

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"regexp"
	"sort"
	"strings"
)

func FetchLabels(patternString, logInf string) map[string]string {
	pattern := regexp.MustCompile(patternString)
	metricsLabelsValues := map[string]string{}
	rslt := pattern.FindStringSubmatch(logInf)
	if rslt == nil {
		return nil
	}
	rslt = rslt[1:]
	groupNames := pattern.SubexpNames()[1:]
	for indx, val := range groupNames {
		metricsLabelsValues[val] = rslt[indx]
	}
	return metricsLabelsValues
}

type Metric struct {
	Gauge   *prometheus.GaugeVec
	Counter *prometheus.CounterVec
	// can add here more metric types
}

var Metrics = map[string]Metric{}

var PromtheusRegistery = prometheus.NewRegistry()

func FetchGroups(rgxPattern string) map[string]Metric {
	pattern := regexp.MustCompile(rgxPattern)
	groupsArray := pattern.SubexpNames()[1:]
	for _, group := range groupsArray {
		if group == "" {
			continue
		}
		groupDetail := strings.Split(group, "_")
		typeValue := groupDetail[0]
		nameValue := groupDetail[1]
		if typeValue == "M" {
			metricType := groupDetail[2]
			switch metricType {
			case "gauge":
				labelsList := []string{}
				for _, lookingLabel := range groupsArray {
					detail := strings.Split(lookingLabel, "_")
					if detail[0] == "L" && detail[1] == nameValue {
						labelsList = append(labelsList, detail[2])
					}
				}
				sort.Strings(labelsList)
				Metrics[nameValue] = Metric{
					Gauge: prometheus.NewGaugeVec(
						prometheus.GaugeOpts{
							Name: nameValue,
						},
						labelsList,
					),
				}
				err := PromtheusRegistery.Register(Metrics[nameValue].Gauge)
				if err != nil {
					log.Fatal("couldn't register metric because: \n", err)
				}
			case "counter":
				labelsList := []string{}
				for _, lookingLabel := range groupsArray {
					detail := strings.Split(lookingLabel, "_")
					if detail[0] == "L" && detail[1] == nameValue {
						labelsList = append(labelsList, detail[2])
					}
				}
				Metrics[nameValue] = Metric{
					Counter: prometheus.NewCounterVec(
						prometheus.CounterOpts{
							Name: nameValue,
						},
						labelsList,
					),
				}
				err := PromtheusRegistery.Register(Metrics[nameValue].Gauge)
				if err != nil {
					log.Fatal("couldn't register metric because: \n", err)
				}
				// can add here additional metric types cases
			}
		}
	}
	return Metrics
}
