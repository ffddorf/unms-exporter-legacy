package main

import (
	"context"
	"fmt"
	"log"

	unms "github.com/ffddorf/unms-api-go"
)

func main() {
	config := unms.NewConfiguration()
	config.BasePath = "http://unms-demo.ubnt.com/v2.0"
	client := unms.NewAPIClient(config)
	auth := context.WithValue(context.Background(), unms.ContextAPIKey, unms.APIKey{
		Key: "123",
	})
	deviceStatusOverview, _, err := client.DevicesApi.DevicesGet(auth, "1", nil)
	if err != nil {
		log.Fatal(err)
	}
	for index, device := range deviceStatusOverview {
		fmt.Println(index, device.Identification.Name, device.Overview.Status)
	}
}
