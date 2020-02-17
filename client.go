package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ffddorf/unms-api-go/models"
)

type DeviceInterfaceSchema struct {

	// addresses
	Addresses models.Addresses `json:"addresses,omitempty"`

	// bridge
	Bridge string `json:"bridge,omitempty"`

	// can display statistics
	CanDisplayStatistics bool `json:"canDisplayStatistics,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// identification
	// Required: true
	Identification *models.InterfaceIdentification `json:"identification"`

	// is switched port
	IsSwitchedPort bool `json:"isSwitchedPort,omitempty"`

	// lag
	Lag string `json:"lag,omitempty"`

	// mtu
	Mtu string `json:"mtu,omitempty"`

	// ospf
	Ospf *models.InterfaceOspf `json:"ospf,omitempty"`

	// poe
	Poe *models.InterfacePoe `json:"poe,omitempty"`

	// pppoe
	Pppoe bool `json:"pppoe,omitempty"`

	// proxy a r p
	ProxyARP bool `json:"proxyARP,omitempty"`

	// speed
	Speed string `json:"speed,omitempty"`

	// statistics
	Statistics *models.InterfaceStatistics `json:"statistics,omitempty"`

	// status
	Status *models.InterfaceStatus `json:"status,omitempty"`

	// switch
	Switch string `json:"switch,omitempty"`

	// visible
	Visible bool `json:"visible,omitempty"`

	// wireless
	Wireless *models.Wireless `json:"wireless,omitempty"`
}

type Device struct {
	Identification struct {
		Hostname string `json:"hostname"`
		ID       string `json:"id"`
	} `json:"identification"`
	Overview struct {
		Status string `json:"status"`
	} `json:"overview"`
	Interfaces []*DeviceInterfaceSchema `json:"interfaces"`
}

type Client struct {
	baseURL string
	token   string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

func (c *Client) DevicesInterfaceStats(siteID string) ([]*Device, error) {
	reqURL := fmt.Sprintf("%s/devices?siteId=%s&withInterfaces=true", c.baseURL, siteID)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-auth-token", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var devices []*Device
	return devices, json.NewDecoder(resp.Body).Decode(&devices)
}
