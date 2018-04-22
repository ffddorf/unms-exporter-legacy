package main

import (
	"context"
	"fmt"
	"log"
	"os"

	unms "github.com/ffddorf/unms-api-go"
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
	deviceStatusOverview, _, err := client.DevicesApi.DevicesGet(auth, "1", nil)
	if err != nil {
		log.Fatal(err)
	}
	for index, device := range deviceStatusOverview {
		fmt.Println(index, device.Identification.Name, device.Overview.Status)
	}
}
