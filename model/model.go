package model

type DeviceRegistration struct {
	Id       string `json:"id"`
	DeviceId string `json:"device_id"`
	Code     string `json:"code"`
}

var deviceRegistration []DeviceRegistration
