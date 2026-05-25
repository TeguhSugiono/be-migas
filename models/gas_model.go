package models

import "time"

type GasLog struct {
	ID int `json:"id"`

	DeviceID string `json:"device_id"`

	StatusGas   bool   `json:"status_gas"`
	AlarmStatus string `json:"alarm_status"`
	WifiStatus  bool   `json:"wifi_status"`

	WASent string `json:"wa_sent"`

	CreatedAt time.Time `json:"created_at"`
}

type GasInput struct {
	DeviceID    string `json:"device_id"`
	DeviceToken string `json:"device_token"`

	StatusGas   bool   `json:"status_gas"`
	AlarmStatus string `json:"alarm_status"`
	WifiStatus  bool   `json:"wifi_status"`
}

// type GasLog struct {
// 	DeviceID    string `json:"device_id"`
// 	DeviceToken string `json:"device_token"`

// 	StatusGas   bool   `json:"status_gas"`
// 	AlarmStatus string `json:"alarm_status"`
// 	WifiStatus  bool   `json:"wifi_status"`
// }

// type GasInput struct {
// 	DeviceID    string `json:"device_id"`
// 	DeviceToken string `json:"device_token"`

// 	StatusGas   bool   `json:"status_gas"`
// 	AlarmStatus string `json:"alarm_status"`
// 	WifiStatus  bool   `json:"wifi_status"`
// }
