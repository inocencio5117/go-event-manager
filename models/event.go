package models

import (
	"database/sql"
	"time"
)

// Event represents an event with various attributes.
// @description Represents an event with various attributes.
type Event struct {
	ID          int          `gorm:"primaryKey" json:"id" swaggertype:"integer"`
	Title       string       `json:"title" swaggertype:"string"`
	Description string       `json:"description" swagertype:"string"`
	StartAt     sql.NullTime `json:"start_at" swaggertype:"string" time_format:"2006-01-02T15:04:05Z"`
	EndAt       sql.NullTime `json:"end_at" swaggertype:"string" time_format:"2006-01-02T15:04:05Z"`
	Location    *string      `json:"location" swaggertype:"string"`
	Status      *string      `json:"status" swaggertype:"string"`
	Capacity    int          `json:"capacity" swaggertype:"integer"`
	CreatedAt   *time.Time   `json:"created_at" swaggerignore:"true" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time   `json:"updated_at" swaggerignore:"true" gorm:"autoUpdateTime"`
}
