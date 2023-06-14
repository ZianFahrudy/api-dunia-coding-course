package service

import (
	"api-dunia-coding/member"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(member member.Member) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("startup_secret_k3y")

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(member member.Member) (string, error) {
	// Tentukan durasi kadaluarsa token
	expirationTime := time.Now().Add(12 * time.Hour) // Misalnya, 1 menit dari sekarang
	claim := jwt.MapClaims{
		"member_id": member.ID,
		"username":  member.Name,
		"exp":       expirationTime.Unix(),
	}
	claim["member_id"] = member.ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil

	})

	if err != nil {
		return token, nil
	}

	return token, nil
}
