package entity

// Device представляет структуру устройства.
type Device struct {
	DeviceID     string `json:"device_id" bson:"device_id"`
	DeviceType   string `json:"deviceType" bson:"deviceType"`
	SerialNumber string `json:"serialNumber" bson:"serialNumber"`
	Status       string `json:"status" bson:"status"`
}
