package main

import (
	"api-dunia-coding/auth"
	"api-dunia-coding/event"
	"api-dunia-coding/handler"
	"api-dunia-coding/helper"
	"api-dunia-coding/member"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:8080)/duniacoding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	memberRepository := member.NewRepository(db)
	eventRepository := event.NewRepository(db)

	memberService := member.NewService(memberRepository)
	authService := auth.NewService()
	eventService := event.NewService(eventRepository)

	memberHandler := handler.NewMemberHandler(memberService, authService)
	eventHandler := handler.NewEventHandler(eventService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	// * Member
	api.POST("/members", memberHandler.RegisterMember)
	api.POST("/login", memberHandler.Login)
	api.POST("/check-email", memberHandler.CheckEmailAvaibility)
	api.POST("/avatar", authMiddleware(authService, memberService), memberHandler.UploadAvatar)

	api.GET("/events", eventHandler.GetEvents)
	api.GET("/events/:id", eventHandler.GetEventDetail)

	router.Run(":8999")

}

func authMiddleware(authService auth.Service, memberService member.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		tokenString := "xx"

		arrayToken := strings.Split(authHeader, " ")

		fmt.Println(len(arrayToken))

		if len(tokenString) == 2 {
			tokenString = arrayToken[1]

		}
		fmt.Println(tokenString)

		token, err := authService.ValidateToken(tokenString)

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
