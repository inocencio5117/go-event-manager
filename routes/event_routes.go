package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/inocencio5117/go-event-manager/handlers"
)

func EventRoutes(r chi.Router) {
	r.Route("/events", func(r chi.Router) {
		r.Post("/", handlers.CreateEvent)
		r.Get("/", handlers.GetAllEvents)
		r.Get("/{id}", handlers.GetEventByID)
		r.Put("/{id}", handlers.UpdateEvent)
		r.Delete("/{id}", handlers.DeleteEvent)

		r.Route("/{event_id}", func(r chi.Router) {
			r.Post("/attendees", handlers.RegisterAttendeeToEvent)
			r.Get("/attendees", handlers.ListAttendeesForEvent)
		})

		r.Route("/{event_id}/notify", func(r chi.Router) {
			r.Post("/", handlers.NotifyAttendees)
		})
	})
}
