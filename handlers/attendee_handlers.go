package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/inocencio5117/go-event-manager/config"
	"github.com/inocencio5117/go-event-manager/models"
)

// @Summary		Register an attendee to an event
// @Description	Registers a new attendee to an existing event.
// @Tags			attendees
// @Accept			json
// @Produce		json
// @Param			attendee	body		models.Attendee	true	"Attendee details"
// @Success		201			{object}	models.Attendee
// @Failure		400			{string}	string	"Invalid request payload"
// @Router			/events/{event_id}/attendees [post]
func RegisterAttendeeToEvent(w http.ResponseWriter, r *http.Request) {
	var attendee models.Attendee

	if err := render.DecodeJSON(r.Body, &attendee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Create(&attendee)
	render.JSON(w, r, attendee)
}

// @Summary		List all attendees for an event
// @Description	List all attendees that where registered in an event
// @Tags			attendees
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.Attendee
// @Failure		400	{string}	string	"Attendees not found"
// @Router			/events/{event_id}/attendees [get]
func ListAttendeesForEvent(w http.ResponseWriter, r *http.Request) {
	var attendees []models.Attendee
	eventId, err := strconv.Atoi(chi.URLParam(r, "event_id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Where("event_id = ?", eventId).Find(&attendees)
	render.JSON(w, r, attendees)
}
