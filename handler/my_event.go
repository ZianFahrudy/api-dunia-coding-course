package handler

import (
	"api-dunia-coding/event"
	"api-dunia-coding/helper"
	"api-dunia-coding/member"
	myevent "api-dunia-coding/my_event"
	"net/http"

	"github.com/gin-gonic/gin"
)

type myEventHandler struct {
	myEventService myevent.Service
}

func NewMyEventHandler(myEventService myevent.Service) *myEventHandler {
	return &myEventHandler{myEventService}
}

func (h *myEventHandler) GetMyEvents(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(member.Member)

	events, err := h.myEventService.GetMyEvents(currentUser.ID)

	if err != nil {
		response := helper.APIResponse("Get Events Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get Events Success", http.StatusOK, event.FormatEvents(events))
	c.JSON(http.StatusOK, response)

}

func (h *myEventHandler) GetUpcomingMyEvents(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(member.Member)

	events, err := h.myEventService.GetUpcomingMyEvents(currentUser.ID)

	if err != nil {
		response := helper.APIResponse("Get Upcoming Events Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get Upcoming Events Success", http.StatusOK, event.FormatEvents(events))
	c.JSON(http.StatusOK, response)

}

func (h *myEventHandler) Presence(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(member.Member)

	var input myevent.PresenceInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Presence this event failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.myEventService.Presence(input, currentUser.ID)

	if err != nil {
		response := helper.APIResponse("Presence this event failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Presence Success", http.StatusOK, true)
	c.JSON(http.StatusOK, response)
}
