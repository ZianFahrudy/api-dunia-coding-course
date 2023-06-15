package controller

import (
	"api-dunia-coding/config"
	"api-dunia-coding/data/formatter"
	"api-dunia-coding/data/model"
	"api-dunia-coding/domain/repository"
	"api-dunia-coding/middleware"
	"api-dunia-coding/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService service.EventService
	repository.AuthRepository
	repository.EventRepository
	config.Config
}

func NewEventController(eventService *service.EventService, authRepository repository.AuthRepository, eventRepository repository.EventRepository, config config.Config) EventController {
	return EventController{EventService: *eventService, AuthRepository: authRepository, EventRepository: eventRepository, Config: config}
}

func (controller *EventController) Route(app *gin.Engine) {
	api := app.Group("/api/v1")
	api.GET("/events", controller.GetEvents)
	api.GET("/events/:id", controller.GetEventDetail)
	api.GET("/events/week", controller.GetEventsOfWeek)
	api.GET("/events/upcoming", controller.GetUpcomingEvents)
	api.GET("/events/calendar", controller.GetCalendarEvents)
	api.POST("/events/join", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.JoinToEvent)

}

func (controller *EventController) GetEvents(c *gin.Context) {

	events, err := controller.EventService.GetEvents(c.Copy())

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Events Failed",
			Data:    nil,
		})
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
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if isToday && now.Before(start) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")

		} else if date.Before(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if date.After(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")
		} else {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Waiting")

		}
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Events Success",
		Data:    formatter.FormatEvents(events),
	})

}

func (controller *EventController) GetEventDetail(c *gin.Context) {
	var input model.GetEventDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Get Event Detail Failed",
			Data:    nil,
		})
		return
	}

	detailEvent, err := controller.EventService.GetEventByID(c.Copy(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Event Detail Failed",
			Data:    nil,
		})
		return
	}

	if detailEvent.ID == 0 {
		c.JSON(http.StatusOK, model.GeneralResponse{
			Code:    http.StatusOK,
			Message: "Get Detail Events Success",
			Data:    nil,
		})
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
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Live")

	} else if isToday && now.After(start) && now.After(end) {
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Missing")

	} else if isToday && now.Before(start) {
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Upcoming")

	} else if date.Before(today) {
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Missing")

	} else if date.After(today) {
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Upcoming")
	} else {
		controller.EventService.UpdateStatusEvent(c.Copy(), detailEvent.ID, "Waiting")

	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Events Success",
		Data:    formatter.FormatEventDetail(detailEvent),
	})
}

func (controller *EventController) GetEventsOfWeek(c *gin.Context) {
	dateStr := ""

	events, err := controller.EventService.GetEventOfWeek(c.Copy())

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Events Week Failed",
			Data:    nil,
		})
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
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if isToday && now.Before(start) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")

		} else if date.Before(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if date.After(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")
		} else {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Waiting")

		}
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Events Week Success",
		Data:    formatter.FormatEvents(events),
	})

}

func (controller *EventController) GetUpcomingEvents(c *gin.Context) {
	dateStr := ""

	events, err := controller.EventService.GetEventByStatus(c.Copy(), "Upcoming")

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Upcoming Events Failed",
			Data:    nil,
		})
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
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Live")

		} else if isToday && now.After(start) && now.After(end) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if isToday && now.Before(start) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")

		} else if date.Before(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Missing")

		} else if date.After(today) {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Upcoming")
		} else {
			controller.EventService.UpdateStatusEvent(c.Copy(), event.ID, "Waiting")

		}
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Upcoming Events Success",
		Data:    formatter.FormatEvents(events),
	})

}

func (controller *EventController) GetCalendarEvents(c *gin.Context) {
	calendarEvents, err := controller.EventService.GetCalendarEvents(c.Copy())
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Calendar Events Failed",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Calendar Events Success",
		Data:    formatter.FormatCalendarEvents(calendarEvents),
	})
}

func (controller *EventController) JoinToEvent(c *gin.Context) {
	var input model.JoinEventInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Join Failed",
			Data:    nil,
		})
		return
	}

	isAvailable, err := controller.EventService.CheckEventMember(c.Copy(), input.EventID, input.MemberID)
	if err != nil {

		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Join Failed",
			Data:    nil,
		})
		return
	}

	if !isAvailable {
		_, err := controller.EventService.JoinToEvent(c.Copy(), input)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, model.GeneralResponse{
				Code:    http.StatusBadRequest,
				Message: "Join Failed",
				Data:    nil,
			})
			return
		}

		data := gin.H{
			"is_join": true,
		}

		c.JSON(http.StatusOK, model.GeneralResponse{
			Code:    http.StatusOK,
			Message: "Join Success",
			Data:    data,
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Anda Sudah Join pada event ini",
		Data:    nil,
	})

}
