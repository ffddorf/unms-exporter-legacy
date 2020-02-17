package main

type InterfaceIdentification struct {

	// Nullable string.
	Description string `json:"description,omitempty"`

	// Computed display name from name and description
	DisplayName string `json:"displayName,omitempty"`

	// mac
	// Pattern: ^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$
	Mac string `json:"mac,omitempty"`

	// Interface name.
	Name string `json:"name,omitempty"`

	// Physical port position.
	Position float64 `json:"position,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

type InterfaceStatistics struct {

	// dropped
	Dropped float64 `json:"dropped,omitempty"`

	// errors
	Errors float64 `json:"errors,omitempty"`

	// rxbytes
	Rxbytes float64 `json:"rxbytes,omitempty"`

	// rxrate
	Rxrate float64 `json:"rxrate,omitempty"`

	// txbytes
	Txbytes float64 `json:"txbytes,omitempty"`

	// txrate
	Txrate float64 `json:"txrate,omitempty"`
}

type InterfaceStatus struct {

	// current speed
	CurrentSpeed string `json:"currentSpeed,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// plugged
	Plugged bool `json:"plugged,omitempty"`

	// speed
	Speed string `json:"speed,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

type DeviceInterfaceSchema struct {
	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// identification
	// Required: true
	Identification *InterfaceIdentification `json:"identification"`

	// statistics
	Statistics *InterfaceStatistics `json:"statistics,omitempty"`

	// status
	Status *InterfaceStatus `json:"status,omitempty"`
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
