package entity

import "time"

type Member struct {
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
