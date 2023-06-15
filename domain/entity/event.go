package entity

import (
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
	Mentor        Mentor
	Status        string
	JoinedEvents  []JoinedEvents
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type JoinedEvents struct {
	ID        int
	EventID   int
	MemberID  int
	Member    Member
	Presence  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CalendarEvent struct {
	Date  string
	Event []Event
}
