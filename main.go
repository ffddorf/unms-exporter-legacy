package main

import (
	"context"
	"log"
	"net/http"
	"os"

	unms "github.com/ffddorf/unms-api-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	url, exists := os.LookupEnv("UNMS_URL")
	if !exists {
		log.Fatal("Please specify the URL of the UNMS API in your environment via UNMS_URL")
	}

	token, exists := os.LookupEnv("UNMS_TOKEN")
	if !exists {
		log.Fatal("Please specify an API_TOKEN to the UNMS API in your environment via UNMS_TOKEN")
	}

	config := unms.NewConfiguration()
	config.BasePath = url
	client := unms.NewAPIClient(config)
	auth := context.WithValue(context.Background(), unms.ContextAPIKey, unms.APIKey{
		Key: token,
	})

	unmsCollector := NewUnmsCollector(client, auth, "")
	prometheus.MustRegister(unmsCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
