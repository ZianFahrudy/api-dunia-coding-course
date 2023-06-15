package main

import (
	"api-dunia-coding/common/exception"
	"api-dunia-coding/config"
	"api-dunia-coding/controller"
	"api-dunia-coding/domain/repository"
	"api-dunia-coding/middleware"
	"api-dunia-coding/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// config
	configuration := config.New()
	db := config.NewDatabase(configuration)

	// repository
	authRepository := repository.NewAuthRepositoryImpl(db)
	eventRepository := repository.NewEventRepositoryImpl(db)
	myEventRepository := repository.NewMyEventRepositoryImpl(db)

	// service
	authService := service.NewAuthServiceImpl(authRepository)
	eventService := service.NewEventServiceImpl(eventRepository)
	myEventService := service.NewMyEventServiceImpl(myEventRepository)

	// controller
	authController := controller.NewAuthController(&authService, authRepository, configuration)
	eventController := controller.NewEventController(&eventService, authRepository, eventRepository, configuration)
	myEventController := controller.NewMyEventController(&myEventService, authRepository, myEventRepository, configuration)

	// Setup Gin
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	app.Static("/storage", "./storage")
	app.Use(gin.CustomRecovery(exception.ErrorHandler))
	app.Use(middleware.CORSMiddleware())

	// Setup Routing
	authController.Route(app)
	eventController.Route(app)
	myEventController.Route(app)

	// Start App
	err := app.Run(configuration.Get("SERVER.PORT"))
	exception.PanicIfNeeded(err)

}
