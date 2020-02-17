package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
