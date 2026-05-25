package models

import "time"

type Device struct {
	ID int `json:"id"`

	DeviceID    string `json:"device_id"`
	DeviceToken string `json:"device_token"`

	DeviceName string `json:"device_name"`

	PairCode string `json:"pair_code"`

	SellerID   *int `json:"seller_id"`
	CustomerID *int `json:"customer_id"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
}

type CreateDeviceInput struct {
	DeviceName string `json:"device_name"`
}

type PairDeviceInput struct {
	DeviceID string `json:"device_id"`
	PairCode string `json:"pair_code"`
}
