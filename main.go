package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"logprom/cmd"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	cmd.Execute()
}
