package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/inocencio5117/go-event-manager/config"
	"github.com/inocencio5117/go-event-manager/models"
)

func NotifyAttendees(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.Atoi(chi.URLParam(r, "event_id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message string
	var attendees []models.Attendee
	var event models.Event

	config.DB.Where("event_id = ?", eventId).Find(&attendees)
	config.DB.Where("id = ?", eventId).Find(&event)

	for _, attendee := range attendees {
		msg := fmt.Sprintf(
			"Event: %s\nDescription: %s\nAttendee: %s\nAttendee email: %s",
			event.Title,
			event.Description,
			attendee.Name,
			attendee.Email,
		)

		if publish := config.PublishNotification("teste", msg); publish.Err() != nil {
			message = fmt.Sprintf("Failed to send notification to %s: %s", attendee.Email, publish.Err().Error())
			break
		}
	}

	if len(attendees) == 0 {
		http.Error(w, "No attendees found for this event", http.StatusBadRequest)
	} else {
		if message != "" {
			log.Printf("Failed to send notification: %s\n", message)
		} else {
			log.Printf("Sent %d notifications for event ID: %d", len(attendees), eventId)

			successMessage := map[string]string{"message": "Notifications sent successfully"}
			render.JSON(w, r, successMessage)
		}
	}
}
