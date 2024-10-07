package model

import (
	"time"
)

// GPSData represents the structure of the data to be saved in the database.
type GPSData struct {
	ID          uint      `gorm:"primaryKey"`     // Auto-incremented ID
	Imei        string    `gorm:"index"`          // Index for quick lookups by IMEI
	TimestampMs uint64    `gorm:"not null"`       // Milliseconds timestamp
	Lng         float64   `gorm:"not null"`       // Longitude
	Lat         float64   `gorm:"not null"`       // Latitude
	Altitude    int16     `gorm:"not null"`       // Altitude
	Angle       uint16    `gorm:"not null"`       // Angle
	EventID     uint16    `gorm:"not null"`       // Event ID
	Speed       uint16    `gorm:"not null"`       // Speed
	Satellites  uint8     `gorm:"not null"`       // Satellites
	Priority    uint8     `gorm:"not null"`       // Priority
	Elements    string    `gorm:"type:jsonb"`     // JSONB type for storing decoded elements
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Auto timestamp
}
