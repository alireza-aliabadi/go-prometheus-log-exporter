package rwfiles

import (
	"github.com/hpcloud/tail"
	"io"
	"log"
	"logprom/internal/prometheus_metrics"
	"logprom/internal/regex_extractor"
)

func ReadFile(pathMetric ...string) {
	path := pathMetric[0]
	metric := pathMetric[1]
	registredMetrics := rgx_extract.FetchGroups(pathMetric[2])
	if metric == "" {
		metric = "log" // default metric
	}
	if len(pathMetric) > 3 {
		log.Fatal("extra parameters are given, only needed prameters are: path, metric")
	}
	t, err := tail.TailFile(path, tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: int64(0),
			Whence: io.SeekStart,
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
			if metrics != nil {
				prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			}
		}
	// add other metrics here as new case
	case "error-count":
		for line := range t.Lines {
			metrics := rgx_extract.FetchLabels(pathMetric[2], line.Text)
			if metrics != nil {
				prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			}
		}

	default:
		for line := range t.Lines {
			metrics := rgx_extract.FetchLabels(pathMetric[2], line.Text)
			if metrics != nil {
				prometheus_metrics.UpdateMetric(metrics, registredMetrics)
			}
		}
	}
}
