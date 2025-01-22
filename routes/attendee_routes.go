package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/inocencio5117/go-event-manager/handlers"
)

func AttendeeRoutes(r chi.Router) {
	r.Route("/events/{event_id}", func(r chi.Router) {
		r.Post("/attendees", handlers.RegisterAttendeeToEvent)
		r.Get("/attendees", handlers.ListAttendeesForEvent)
	})
}
