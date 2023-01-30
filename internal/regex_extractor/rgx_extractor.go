package rgx_extract

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"regexp"
	"strings"
)

func FetchLabels(patternString, logInf string) map[string]string {
	pattern := regexp.MustCompile(patternString)
	//for _, name := range pattern.SubexpNames()[1:] {
	//	fmt.Println("type of elements --> ", string(name[0]))
	//}
	rslt := pattern.FindStringSubmatch(logInf)[1:]
	groupNames := pattern.SubexpNames()[1:]
	metricsLabelsValues := map[string]string{}
	for indx, val := range groupNames {
		metricsLabelsValues[val] = rslt[indx]
	}
	fmt.Println("result ---> \n", metricsLabelsValues)
	return metricsLabelsValues
}

type Metric struct {
	Gauge   *prometheus.GaugeVec
	Counter *prometheus.CounterVec
}

var Metrics = map[string]Metric{}

var PromtheusRegistery = prometheus.NewRegistry()

func FetchGroups(rgxPattern string) map[string]Metric {
	pattern := regexp.MustCompile(rgxPattern)
	groupsArray := pattern.SubexpNames()[1:]
	for _, group := range groupsArray {
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
			}
		}
	}
	fmt.Println("result metrics --> \n", Metrics)
	return Metrics
}
