package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	unmsURL, exists := os.LookupEnv("UNMS_URL")
	if !exists {
		log.Fatal("Please specify the URL of the UNMS API in your environment via UNMS_URL")
	}

	token, exists := os.LookupEnv("UNMS_TOKEN")
	if !exists {
		log.Fatal("Please specify an API_TOKEN to the UNMS API in your environment via UNMS_TOKEN")
	}

	siteID, exists := os.LookupEnv("UNMS_SITE_ID")
	if !exists {
		log.Fatal("Please specify a Site ID in your environment via UNMS_SITE_ID")
	}

	client := NewClient(unmsURL, token)
	collector := NewUnmsCollector(client, siteID)
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
