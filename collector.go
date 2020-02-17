package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type unmsCollector struct {
	client *Client
	siteID string

	metricDeviceUp         *prometheus.Desc
	metricInterfaceUp      *prometheus.Desc
	metricInterfaceRXBytes *prometheus.Desc
	metricInterfaceTXBytes *prometheus.Desc
}

const namespace = "unms"

func NewUnmsCollector(client *Client, siteID string) prometheus.Collector {
	c := unmsCollector{
		client: client,
		siteID: siteID,
		metricDeviceUp: prometheus.NewDesc(
			namespace+"_device_up",
			"If device is up",
			[]string{"device_id", "hostname"},
			nil,
		),
		metricInterfaceUp: prometheus.NewDesc(
			namespace+"_interface_up",
			"If interface is up",
			[]string{"device_id", "hostname", "interface"},
			nil,
		),
		metricInterfaceRXBytes: prometheus.NewDesc(
			namespace+"_interface_rx_bytes",
			"Bytes received",
			[]string{"device_id", "hostname", "interface"},
			nil,
		),
		metricInterfaceTXBytes: prometheus.NewDesc(
			namespace+"_interface_tx_bytes",
			"Bytes transmitted",
			[]string{"device_id", "hostname", "interface"},
			nil,
		),
	}

	return &c
}

func (c *unmsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metricDeviceUp
	ch <- c.metricInterfaceUp
	ch <- c.metricInterfaceRXBytes
	ch <- c.metricInterfaceTXBytes
}

func (c *unmsCollector) collectMetricDevice(device *Device) prometheus.Metric {
	deviceUp := float64(0)
	if device.Overview.Status == "active" {
		deviceUp = 1
	}
	return prometheus.MustNewConstMetric(
		c.metricDeviceUp,
		prometheus.GaugeValue,
		deviceUp,
		device.Identification.ID,
		device.Identification.Hostname,
	)
}

func (c *unmsCollector) collectMetricInterface(ch chan<- prometheus.Metric, device *Device) {
	for _, intf := range device.Interfaces {
		intfUp := float64(0)
		if intf.Status.Status == "active" {
			intfUp = 1
		}
		ch <- prometheus.MustNewConstMetric(
			c.metricInterfaceUp,
			prometheus.GaugeValue,
			intfUp,
			device.Identification.ID,
			device.Identification.Hostname,
			intf.Identification.DisplayName,
		)

		if intf.Statistics == nil {
			continue
		}

		rxBytes := intf.Statistics.Rxbytes
		ch <- prometheus.MustNewConstMetric(
			c.metricInterfaceRXBytes,
			prometheus.CounterValue,
			rxBytes,
			device.Identification.ID,
			device.Identification.Hostname,
			intf.Identification.DisplayName,
		)

		txBytes := intf.Statistics.Txbytes
		ch <- prometheus.MustNewConstMetric(
			c.metricInterfaceTXBytes,
			prometheus.CounterValue,
			txBytes,
			device.Identification.ID,
			device.Identification.Hostname,
			intf.Identification.DisplayName,
		)
	}
}

func (c *unmsCollector) Collect(ch chan<- prometheus.Metric) {
	devices, err := c.client.DevicesInterfaceStats(c.siteID)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	for _, dev := range devices {
		ch <- c.collectMetricDevice(dev)
		c.collectMetricInterface(ch, dev)
	}
}
