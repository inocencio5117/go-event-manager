package models

import "time"

// Attendee represents a participant in an event.
// @description Attendee model represents a participant in an event.
type Attendee struct {
	ID        uint       `gorm:"primaryKey" json:"id" swaggertype:"integer"`
	Name      string     `json:"name" swaggertype:"string" validate:"required,min=2,max=125"`
	Email     string     `json:"email" swaggertype:"string" validate:"required,email"`
	EventID   uint       `json:"event_id" swaggertype:"integer" gorm:"index:idx_event_attendee"`
	CreatedAt *time.Time `json:"created_at" swaggerignore:"true" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" swaggerignore:"true" gorm:"autoUpdateTime"`
}
