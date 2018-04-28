package main

import (
	"context"
	"log"

	"github.com/ffddorf/unms-api-go"
	"github.com/prometheus/client_golang/prometheus"
)

type unmsCollector struct {
	client    *unms.APIClient
	clientCtx context.Context
	site      string
	up        *prometheus.Desc
}

const namespace = "unms"

func NewUnmsCollector(client *unms.APIClient, clientCtx context.Context, site string) prometheus.Collector {
	c := unmsCollector{
		client:    client,
		clientCtx: clientCtx,
		site:      site,
		up: prometheus.NewDesc(
			namespace+"_"+"device_up",
			"If device is connected to UNMS.",
			[]string{"id"},
			nil,
		),
	}

	return &c
}

func (c *unmsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
}

func (c *unmsCollector) Collect(ch chan<- prometheus.Metric) {
	deviceStatusOverview, _, err := c.client.DevicesApi.DevicesGet(c.clientCtx, c.site, nil)
	if err != nil {
		log.Println(err)
	}
	for _, device := range deviceStatusOverview {
		up := float64(0)
		if device.Overview.Status == "active" {
			up = 1
		}
		ch <- prometheus.MustNewConstMetric(
			c.up,
			prometheus.GaugeValue,
			up,
			device.Identification.Id,
		)
	}
}
