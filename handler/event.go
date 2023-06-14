package handler

import (
	"api-dunia-coding/event"
	"api-dunia-coding/helper"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventService event.Service
}

func NewEventHandler(eventService event.Service) *eventHandler {
	return &eventHandler{eventService}
}

func (h *eventHandler) GetEvents(c *gin.Context) {

	events, err := h.eventService.GetEvents()

	if err != nil {
		response := helper.APIResponse("Get Events Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, event := range events {

		layout := "2006-01-02"

		date, err := time.Parse(layout, event.Date)
		if err != nil {
			panic(err)
		}

		today := time.Now()
		isToday := date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()

		startTime := event.StartTime
		endTime := event.EndTime

		partsStart := strings.Split(startTime, ":")
		startHour := partsStart[0]
		numStart, _ := strconv.Atoi(startHour)

		partsEnd := strings.Split(endTime, ":")
		endHour := partsEnd[0]

		numEnd, _ := strconv.Atoi(endHour)
		now := time.Now()

		start := time.Date(now.Year(), now.Month(), now.Day(), numStart, 0, 0, 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), numEnd, 0, 0, 0, now.Location())
		if isToday && now.After(start) && now.Before(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if isToday && now.Before(start) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")

		} else if date.Before(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if date.After(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")
		} else {
			h.eventService.UpdateStatusEvent(event.ID, "Waiting")

		}
	}

	response := helper.APIResponse("Get Events Success", http.StatusOK, event.FormatEvents(events))
	c.JSON(http.StatusOK, response)

}

func (h *eventHandler) GetEventDetail(c *gin.Context) {
	var input event.GetEventDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Get Event Detail Failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	detailEvent, err := h.eventService.GetEventByID(input)
	if err != nil {
		response := helper.APIResponse("Get Event Detail Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if detailEvent.ID == 0 {
		response := helper.APIResponse("Get Detail Event Success", http.StatusOK, nil)
		c.JSON(http.StatusOK, response)
		return
	}
	layout := "2006-01-02"
	dateString := detailEvent.Date

	date, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}

	today := time.Now()
	isToday := date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()

	startTime := detailEvent.StartTime
	endTime := detailEvent.EndTime

	partsStart := strings.Split(startTime, ":")
	startHour := partsStart[0]
	numStart, _ := strconv.Atoi(startHour)

	partsEnd := strings.Split(endTime, ":")
	endHour := partsEnd[0]

	numEnd, _ := strconv.Atoi(endHour)
	now := time.Now()

	start := time.Date(now.Year(), now.Month(), now.Day(), numStart, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), numEnd, 0, 0, 0, now.Location())
	if isToday && now.After(start) && now.Before(end) {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Live")

	} else if isToday && now.After(start) && now.After(end) {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Missing")

	} else if isToday && now.Before(start) {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Upcoming")

	} else if date.Before(today) {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Missing")

	} else if date.After(today) {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Upcoming")
	} else {
		h.eventService.UpdateStatusEvent(detailEvent.ID, "Waiting")

	}
	response := helper.APIResponse("Get Detail Event Success", http.StatusOK, event.FormatEventDetail(detailEvent))
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) GetEventsOfWeek(c *gin.Context) {
	dateStr := ""

	events, err := h.eventService.GetEventOfWeek()

	if err != nil {
		response := helper.APIResponse("Get Events Week Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, event := range events {
		layout := "2006-01-02"

		if dateStr == "" {
			dateStr = event.Date
		}

		date, err := time.Parse(layout, dateStr)
		if err != nil {
			panic(err)
		}

		today := time.Now()
		isToday := date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()

		startTime := event.StartTime
		endTime := event.EndTime

		partsStart := strings.Split(startTime, ":")
		startHour := partsStart[0]
		numStart, _ := strconv.Atoi(startHour)

		partsEnd := strings.Split(endTime, ":")
		endHour := partsEnd[0]

		numEnd, _ := strconv.Atoi(endHour)
		now := time.Now()

		start := time.Date(now.Year(), now.Month(), now.Day(), numStart, 0, 0, 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), numEnd, 0, 0, 0, now.Location())
		if isToday && now.After(start) && now.Before(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if isToday && now.Before(start) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")

		} else if date.Before(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if date.After(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")
		} else {
			h.eventService.UpdateStatusEvent(event.ID, "Waiting")

		}
	}

	response := helper.APIResponse("Get Events Week Success", http.StatusOK, event.FormatEvents(events))
	c.JSON(http.StatusOK, response)

}

func (h *eventHandler) GetUpcomingEvents(c *gin.Context) {
	dateStr := ""

	events, err := h.eventService.GetEventByStatus("Upcoming")

	if err != nil {
		response := helper.APIResponse("Get Upcoming Events Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, event := range events {
		layout := "2006-01-02"

		if dateStr == "" {
			dateStr = event.Date
		}

		date, err := time.Parse(layout, dateStr)
		if err != nil {
			panic(err)
		}

		today := time.Now()
		isToday := date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()

		startTime := event.StartTime
		endTime := event.EndTime

		partsStart := strings.Split(startTime, ":")
		startHour := partsStart[0]
		numStart, _ := strconv.Atoi(startHour)

		partsEnd := strings.Split(endTime, ":")
		endHour := partsEnd[0]

		numEnd, _ := strconv.Atoi(endHour)
		now := time.Now()

		start := time.Date(now.Year(), now.Month(), now.Day(), numStart, 0, 0, 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), numEnd, 0, 0, 0, now.Location())
		if isToday && now.After(start) && now.Before(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if isToday && now.Before(start) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")

		} else if date.Before(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Missing")

		} else if date.After(today) {
			h.eventService.UpdateStatusEvent(event.ID, "Upcoming")
		} else {
			h.eventService.UpdateStatusEvent(event.ID, "Waiting")

		}
	}

	response := helper.APIResponse("Get Upcoming Events Success", http.StatusOK, event.FormatEvents(events))
	c.JSON(http.StatusOK, response)

}

func (h *eventHandler) GetCalendarEvents(c *gin.Context) {
	calendarEvents, err := h.eventService.GetCalendarEvents()
	if err != nil {
		response := helper.APIResponse("Get Calendar Events Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Get Upcoming Events Success", http.StatusOK, event.FormatCalendarEvents(calendarEvents))
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) JoinToEvent(c *gin.Context) {
	var input event.JoinEventInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Join Failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := h.eventService.CheckEventMember(input.EventID, input.MemberID)
	if err != nil {
		response := helper.APIResponse("Join Failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isAvailable {
		_, err := h.eventService.JoinToEvent(input)
		if err != nil {
			response := helper.APIResponse("Join Failed", http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		data := gin.H{
			"is_join": true,
		}

		response := helper.APIResponse("Join Success", http.StatusOK, data)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Anda Sudah Join pada event ini", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)

}
