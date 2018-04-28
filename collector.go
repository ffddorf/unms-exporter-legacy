package main

import (
	"context"
	"log"

	"github.com/ffddorf/unms-api-go"
	"github.com/prometheus/client_golang/prometheus"
)

type unmsCollector struct {
	client *unms.APIClient
	clientCtx context.Context
	site string
	devices []unms.DeviceStatusOverview
	metric_device_name *prometheus.Desc
	metric_device_up *prometheus.Desc
}

const namespace = "unms"

func NewUnmsCollector(client *unms.APIClient, clientCtx context.Context, site string) prometheus.Collector {
	c := unmsCollector{
		client: client,
		clientCtx: clientCtx,
		site: site,
		devices: []unms.DeviceStatusOverview{},
		metric_device_name: prometheus.NewDesc(
			namespace+"_"+"device_name",
			"The ID and name of a device. Value is always 0.",
			[]string{"id","name"},
			nil,
		),
		metric_device_up: prometheus.NewDesc(
			namespace+"_"+"device_up",
			"If device is connected to UNMS.",
			[]string{"id"},
			nil,
		),
	}

	return &c
}


func (c *unmsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metric_device_name
	ch <- c.metric_device_up
}

func (c *unmsCollector) getDeviceStatusOverview() {
	deviceStatusOverview, _, err := c.client.DevicesApi.DevicesGet(c.clientCtx, c.site, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c.devices = deviceStatusOverview
}

func (c *unmsCollector) collectMetricDeviceName(ch chan<- prometheus.Metric) {
	for _, device := range c.devices {
		ch <- prometheus.MustNewConstMetric (
			c.metric_device_name,
			prometheus.GaugeValue,
			float64(0),
			device.Identification.Id,
			device.Identification.Name,
		)
	}
}

func (c *unmsCollector) collectMetricDeviceUp(ch chan<- prometheus.Metric) {
	for _, device := range c.devices {
		device_up := float64(0)
		if device.Overview.Status == "active" {
			device_up = 1
		}
		ch <- prometheus.MustNewConstMetric(
			c.metric_device_up,
			prometheus.GaugeValue,
			device_up,
			device.Identification.Id,
		)
	}
}

func (c *unmsCollector) Collect(ch chan<- prometheus.Metric) {
	c.getDeviceStatusOverview()
	c.collectMetricDeviceUp(ch)
	c.collectMetricDeviceName(ch)
}
