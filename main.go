package main

import (
	// "api-dunia-coding/auth"

	"api-dunia-coding/config"
	"api-dunia-coding/controller"
	"api-dunia-coding/exception"
	"api-dunia-coding/helper"
	"api-dunia-coding/member"
	"api-dunia-coding/middleware"
	"api-dunia-coding/repository"
	"api-dunia-coding/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	configuration := config.New()
	db := config.NewDatabase(configuration)

	//! repository

	authRepository := repository.NewAuthRepositoryImpl(db)

	authService := service.NewAuthServiceImpl(authRepository)

	authController := controller.NewAuthController(&authService, authRepository, configuration)

	// memberRepository := member.NewRepository(db)
	// eventRepository := event.NewRepository(db)
	// myEventRepository := myevent.NewRepository(db)

	// //! service
	// memberService := member.NewService(memberRepository)
	// authService := auth.NewService()
	// eventService := event.NewService(eventRepository)
	// myEventService := myevent.NewService(myEventRepository)

	// //! handler
	// memberHandler := handler.NewMemberHandler(memberService, authService)
	// eventHandler := handler.NewEventHandler(eventService)
	// myEventHandler := handler.NewMyEventHandler(myEventService)

	// router := gin.Default()
	// router.Static("/images", "./images")
	// api := router.Group("/api/v1")

	// // *Member
	// api.POST("/members", memberHandler.RegisterMember)
	// api.POST("/login", memberHandler.Login)
	// api.POST("/check-email", memberHandler.CheckEmailAvaibility)
	// api.POST("/avatar", authMiddleware(authService, memberService), memberHandler.UploadAvatar)

	// // * Event
	// api.GET("/events", eventHandler.GetEvents)
	// api.GET("/events/:id", eventHandler.GetEventDetail)
	// api.GET("/events/week", eventHandler.GetEventsOfWeek)
	// api.GET("/events/upcoming", eventHandler.GetUpcomingEvents)
	// api.GET("/events/calendar", eventHandler.GetCalendarEvents)
	// api.POST("/events/join", authMiddleware(authService, memberService), eventHandler.JoinToEvent)

	// // * My Event
	// api.GET("/myevents", authMiddleware(authService, memberService), myEventHandler.GetMyEvents)
	// api.GET("/myevents/upcoming", authMiddleware(authService, memberService), myEventHandler.GetUpcomingMyEvents)
	// api.POST("/myevents/presence/:id", authMiddleware(authService, memberService), myEventHandler.Presence)

	// router.Run(":8081")

	// Setup Gin
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	app.Static("/images", "./images")
	app.Use(gin.CustomRecovery(exception.ErrorHandler))
	app.Use(middleware.CORSMiddleware())

	// Setup Routing
	authController.Route(app)

	// Start App
	err := app.Run(configuration.Get("SERVER.PORT"))
	exception.PanicIfNeeded(err)

}

func authMiddleware(jwtService service.JwtService, memberService member.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		tokenString := "xx"

		arrayToken := strings.Split(authHeader, " ")

		// fmt.Println(len(arrayToken))

		if len(tokenString) == 2 {
			tokenString = arrayToken[1]

		}
		// fmt.Println(tokenString)

		token, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		memberID := int(claim["member_id"].(float64))

		member, err := memberService.GetMemberByID(memberID)

		if err != nil {
			response := helper.APIResponse("Unauthotized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		c.Set("currentUser", member)

	}

}
