package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/inocencio5117/go-event-manager/config"
	"github.com/inocencio5117/go-event-manager/models"
)

// @Summary		Create a new event
// @Description	Create a new event in the database
// @Tags			events
// @Accept			json
// @Produce		json
// @Param			event	body		models.Event	true	"Event details"
// @Success		201		{object}	models.Event
// @Failure		400		{string}	string	"Invalid request payload"
// @Router			/events [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	if err := render.DecodeJSON(r.Body, &event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.DB.Create(&event)
	render.JSON(w, r, event)
}

// @Summary		Get all events
// @Description	Get all events from the database
// @Tags			events
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.Event
// @Failure		500	{string}	string	"Internal server error"
// @Router			/events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event

	config.DB.Find(&events)
	render.JSON(w, r, events)
}

// @Summary		Get an event by ID
// @Description	Retrieve a single event by its ID
// @Tags			events
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Event ID"
// @Success		200	{object}	models.Event
// @Failure		404	{string}	string	"Event not found"
// @Router			/events/{id} [get]
func GetEventByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var event models.Event

	if err := config.DB.First(&event, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	render.JSON(w, r, event)
}

// @Summary		Update an event
// @Description	Update an existing event by its ID
// @Tags			events
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Event ID"
// @Param			event	body		models.Event	true	"Event details to update"
// @Success		200		{object}	models.Event
// @Failure		404		{string}	string	"Event not found"
// @Router			/events/{id} [put]
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var event models.Event

	if err := render.DecodeJSON(r.Body, &event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := render.DecodeJSON(r.Body, id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	config.DB.Save(&event)
	render.JSON(w, r, event)
}

// @Summary		Delete an event
// @Description	Delete an existing event by its ID
// @Tags			events
// @Param			id	path		int		true	"Event ID"
// @Success		204	{string}	string	"Event deleted"
// @Failure		404	{string}	string	"Event not found"
// @Router			/events/{id} [delete]
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var event models.Event

	if err := config.DB.First(&event, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	config.DB.Delete(&event)
	w.WriteHeader(http.StatusOK)
}
