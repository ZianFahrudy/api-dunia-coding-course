package controller

import (
	"api-dunia-coding/config"
	"api-dunia-coding/data/formatter"
	"api-dunia-coding/data/model"
	"api-dunia-coding/domain/repository"
	"api-dunia-coding/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InformationController struct {
	InformationService service.InformationService
	repository.AuthRepository
	repository.InformationRepository
	config.Config
}

func NewInformationController(service *service.InformationService, authRepository repository.AuthRepository, informationRepository repository.InformationRepository, config config.Config) InformationController {
	return InformationController{InformationService: *service, AuthRepository: authRepository, InformationRepository: informationRepository, Config: config}
}

func (controller *InformationController) Route(app *gin.Engine) {
	api := app.Group("/api/v1")

	api.GET("/informations", controller.GetInformations)
	api.DELETE("/informations/:id", controller.DeteleInformation)
	// api.POST("/informations", controller.DeteleAllInformation)

}

func (controller *InformationController) GetInformations(c *gin.Context) {

	informations, err := controller.InformationService.GetAll(c.Copy())

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed",
			Data:    nil,
		})
		return

	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    formatter.Informations(informations),
	})

}

func (controller *InformationController) DeteleInformation(c *gin.Context) {
	var input model.DeleteInformationInput
	err := c.ShouldBindUri(&input)

	information, err := controller.InformationService.GetByID(c.Copy(), input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed",
			Data:    nil,
		})
		return

	}

	isDeleted, err := controller.InformationService.Delete(c.Copy(), information.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed",
			Data:    nil,
		})
		return

	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    isDeleted,
	})

}

func (controller *InformationController) DeteleAllInformation(c *gin.Context) {
	var input model.DeleteAllInformationInput
	err := c.ShouldBindJSON(&input)

	isDeleted, err := controller.InformationService.DeleteAll(c.Copy(), input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed",
			Data:    nil,
		})
		return

	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    isDeleted,
	})

}
