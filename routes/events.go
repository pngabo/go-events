package routes

import (
	"events-mgmt-portal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. try again later",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	ctx.JSON(http.StatusOK, event)
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event. try again later",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created.", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
