package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	serverPort := ":8080"
	metricsPath := "/metrics"

	http.Handle(metricsPath, promhttp.Handler())

	log.Println("Serving metrics on " + serverPort + metricsPath)
	http.ListenAndServe(serverPort, nil)
}
