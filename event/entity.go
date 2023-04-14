package event

import (
	"api-dunia-coding/member"
	"api-dunia-coding/mentor"
	"time"
)

type Event struct {
	ID            int
	Label         string
	Name          string
	Date          string
	StartTime     string
	EndTime       string
	MeetURL       string
	Platform      string
	About         string
	Description   string
	Documentation string
	MentorID      int
	Mentor        mentor.Mentor
	Status        string
	JoinedEvent   []JoinedEvent
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type JoinedEvent struct {
	ID        int
	EventID   int
	MemberID  int
	Member    member.Member
	CreatedAt time.Time
	UpdatedAt time.Time
}
