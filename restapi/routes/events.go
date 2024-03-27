package routes

import (
	"net/http"
	"strconv"

	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event."})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request event."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event."})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}
	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event delete successfully!"})
}
