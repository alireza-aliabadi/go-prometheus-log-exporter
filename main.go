package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"logprom/cmd"
	"logprom/internal/regex_extractor"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.HandlerFor(rgx_extract.PromtheusRegistery, promhttp.HandlerOpts{}))
	cmd.Execute()
}
