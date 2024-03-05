package entity

type Station struct {
	ID           int    `json:"id"`
	Cabinet      string `json:"cabinet"`
	DeviceNumber string `json:"device_number"`
}
