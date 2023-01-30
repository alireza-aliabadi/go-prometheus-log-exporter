package rwfiles

import (
	"github.com/hpcloud/tail"
	"io"
	"log"
	//logmetric "logprom/internal/logmetrics"
	"logprom/internal/prometheus_metrics"
	"logprom/internal/regex_extractor"
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

func ReadFile(pathMetric ...string) {
	path := pathMetric[0]
	metric := pathMetric[1]
	registredMetrics := rgx_extract.FetchGroups(pathMetric[2])
	if metric == "" {
		metric = "login" // default metric
	}
	if len(pathMetric) > 3 {
		log.Fatal("extra parameters are given, only needed prameters are: path, metric")
	}
	t, err := tail.TailFile(path, tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: int64(0),
			Whence: io.SeekEnd,
		},
		ReOpen: true,
	})
	if err != nil {
		log.Fatal("file tail error: -->", err)
	}
	switch metric {
	case "log":
		for line := range t.Lines {
			metrics := rgx_extract.FetchLabels(pathMetric[2], line.Text)
			prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			//metricDetail := GetLogInf(line.Text)
			//logmetric.LogGaugeVec(metricDetail)
			//logmetric.ResponseTimeGauge(metricDetail)
		}
	// add other metrics here as new case
	case "error-count":
		for line := range t.Lines {
			metrics := rgx_extract.FetchLabels(pathMetric[2], line.Text)
			prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			//metricDetail := GetLogInf(line.Text)
			//logmetric.ErrCounterVec(metricDetail)
		}

	default:
		for line := range t.Lines {
			metrics := rgx_extract.FetchLabels(pathMetric[2], line.Text)
			prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			//metricDetail := GetLogInf(line.Text)
			//logmetric.LogGaugeVec(metricDetail)
		}
	}
}
