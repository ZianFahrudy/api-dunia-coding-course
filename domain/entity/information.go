package entity

import "time"

type Information struct {
	ID        int
	Title     string
	Body      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
