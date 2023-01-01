package rwfiles

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log"
	logmetric "logprom/internal/logmetrics"
	"strings"
)

func GetLogInf(log string) map[string]string {
	resultMap := map[string]string{}
	rslt := strings.Split(log, "|")
	resultMap["log time"] = rslt[0]
	for _, val := range rslt[3:] {
		tempRslt := strings.Split(val, ":")
		resultMap[tempRslt[0]] = tempRslt[1]
	}
	return resultMap
}

func ReadFile(path_metric ...string) {
	path := path_metric[0]
	metric := path_metric[1]
	if len(path_metric) > 2 {
		log.Fatal("extra parameters are given, only needed prameters are: path, metric")
	}
	t, err := tail.TailFile(path, tail.Config{
		Follow: true,
		ReOpen: true})
	if err != nil {
		log.Fatal("file tail error: -->", err)
	}
	switch metric {
	case "login":
		for line := range t.Lines {
			fmt.Println(line.Text)
			metricDetail := GetLogInf(line.Text)
			logmetric.LogGaugeVec(metricDetail)
		}
	// add other metrics here as new case
	case "error-count":
		for line := range t.Lines {
			metricDetail := GetLogInf(line.Text)
			logmetric.ErrCounterVec(metricDetail)
		}

	default:
		for line := range t.Lines {
			fmt.Println(line.Text)
			metricDetail := GetLogInf(line.Text)
			logmetric.LogGaugeVec(metricDetail)
		}
	}
}
