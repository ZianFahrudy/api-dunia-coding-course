package common

import (
	"api-dunia-coding/config"
	"api-dunia-coding/exception"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, memberID int, config config.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicIfNeeded(err)
	claims := jwt.MapClaims{
		"username":  username,
		"member_id": memberID,
		"exp":       time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicIfNeeded(err)

	return tokenSigned
}
