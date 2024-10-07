package model

import "time"

type DataCodec struct {
	ID        uint      `gorm:"primaryKey"` // Auto-incremented ID
	Imei      string    `gorm:"index"`      // IMEI for identifying the device
	TotalData int       // Total data count for the device
	GPSData   []GPSData `gorm:"foreignKey:DataCodecID"` // Relation to GPSData
	CreatedAt time.Time `gorm:"autoCreateTime"`         // Auto timestamp for record creation
}
