package mentor

import "time"

type Mentor struct {
	ID         int
	Name       string
	Occupation string
	AvatarURL  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
