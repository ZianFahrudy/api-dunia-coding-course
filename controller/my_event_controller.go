package controller

import (
	"api-dunia-coding/config"
	"api-dunia-coding/data/formatter"
	"api-dunia-coding/data/model"
	"api-dunia-coding/domain/entity"
	"api-dunia-coding/domain/repository"
	"api-dunia-coding/middleware"
	"api-dunia-coding/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyEventController struct {
	MyEventService service.MyEventService
	repository.AuthRepository
	repository.MyEventRepository
	config.Config
}

func NewMyEventController(myEventService *service.MyEventService, authRepository repository.AuthRepository, myEventRepository repository.MyEventRepository, config config.Config) MyEventController {
	return MyEventController{MyEventService: *myEventService, AuthRepository: authRepository, MyEventRepository: myEventRepository, Config: config}
}

func (controller *MyEventController) Route(app *gin.Engine) {
	api := app.Group("/api/v1")

	api.GET("/myevents", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.GetMyEvents)
	api.GET("/myevents/upcoming", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.GetUpcomingMyEvents)
	api.POST("/myevents/presence/:id", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.Presence)

}

func (controller *MyEventController) GetMyEvents(c *gin.Context) {

	currentUser := c.MustGet(controller.Get("JWT_CURRENT_USER")).(entity.Member)

	events, err := controller.MyEventService.GetMyEvents(c.Copy(), currentUser.ID)

	if err != nil {

		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Events Failed",
			Data:    nil,
		})
		return

	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get My Events Success",
		Data:    formatter.FormatEvents(events),
	})

}

func (controller *MyEventController) GetUpcomingMyEvents(c *gin.Context) {

	currentUser := c.MustGet(controller.Get("JWT_CURRENT_USER")).(entity.Member)

	events, err := controller.MyEventService.GetUpcomingMyEvents(c.Copy(), currentUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Upcoming Events Failed",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Get Upcoming Events Success",
		Data:    events,
	})

}

func (controller *MyEventController) Presence(c *gin.Context) {
	currentUser := c.MustGet(controller.Get("JWT_CURRENT_USER")).(entity.Member)

	var input model.PresenceInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Presence Failed",
			Data:    nil,
		})
		return
	}

	err = controller.MyEventService.Presence(c.Copy(), input, currentUser.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Presence Failed",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Presence Success",
		Data:    true,
	})
}
