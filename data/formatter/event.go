package formatter

import (
	"api-dunia-coding/domain/entity"
)

// response formatter
type EventFormatter struct {
	ID           int                    `json:"id"`
	Label        string                 `json:"label"`
	Name         string                 `json:"name"`
	Date         string                 `json:"date"`
	StartTime    string                 `json:"start_time"`
	EndTime      string                 `json:"end_time"`
	MeetURL      string                 `json:"meet_url"`
	Mentor       MentorFormatter        `json:"mentor"`
	Status       string                 `json:"status"`
	JoinedEvents []JoinedEventFormatter `json:"joined_members"`
}

type EventDetailFormatter struct {
	ID            int             `json:"id"`
	Label         string          `json:"label"`
	Name          string          `json:"name"`
	Date          string          `json:"date"`
	StartTime     string          `json:"start_time"`
	EndTime       string          `json:"end_time"`
	MeetURL       string          `json:"meet_url"`
	Platform      string          `json:"platform"`
	About         string          `json:"about"`
	Description   string          `json:"description"`
	Documentation string          `json:"documentation"`
	Mentor        MentorFormatter `json:"mentor"`
	Status        string          `json:"status"`
}

type JoinedEventFormatter struct {
	Member MemberEventFormatter `json:"member"`
}

type MentorFormatter struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	AvatarURL  string `json:"avatar_url"`
}

type MemberEventFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type CalendarEventFormatter struct {
	Date   string                   `json:"date"`
	Events []EventCalendarFormatter `json:"events"`
}

type EventCalendarFormatter struct {
	ID        int             `json:"id"`
	Label     string          `json:"label"`
	Name      string          `json:"name"`
	Date      string          `json:"date"`
	StartTime string          `json:"start_time"`
	EndTime   string          `json:"end_time"`
	Mentor    MentorFormatter `json:"mentor"`
	Status    string          `json:"status"`
}

// formatter function
func FormatCalendarEvent(event entity.CalendarEvent) CalendarEventFormatter {
	calendarEventFormatter := CalendarEventFormatter{
		Date: event.Date,
	}

	eventsCalendar := []EventCalendarFormatter{}

	for _, event := range event.Event {
		eventCalendarFormatter := EventCalendarFormatter{
			ID:        event.ID,
			Date:      event.Date,
			Label:     event.Label,
			Name:      event.Name,
			StartTime: event.StartTime,
			EndTime:   event.EndTime,
			Status:    event.Status,
			Mentor: MentorFormatter{
				Name:       event.Mentor.Name,
				Occupation: event.Mentor.Occupation,
				AvatarURL:  event.Mentor.AvatarURL,
			},
		}

		calendarEventFormatter.Events = eventsCalendar
		eventsCalendar = append(eventsCalendar, eventCalendarFormatter)
	}
	calendarEventFormatter.Events = eventsCalendar

	return calendarEventFormatter
}

func FormatEvent(event entity.Event) EventFormatter {
	eventFormatter := EventFormatter{
		ID:        event.ID,
		Label:     event.Label,
		Name:      event.Name,
		Date:      event.Date,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		MeetURL:   event.MeetURL,
		Status:    event.Status,
		Mentor: MentorFormatter{
			Name:       event.Mentor.Name,
			Occupation: event.Mentor.Occupation,
			AvatarURL:  event.Mentor.AvatarURL,
		},
	}

	joinedEvents := []JoinedEventFormatter{}

	for _, joinedEvent := range event.JoinedEvents {
		memberEventFormatter := MemberEventFormatter{
			Name:  joinedEvent.Member.Name,
			Email: joinedEvent.Member.Email,
		}
		joinedEventFormatter := JoinedEventFormatter{
			Member: memberEventFormatter,
		}

		joinedEvents = append(joinedEvents, joinedEventFormatter)
	}

	eventFormatter.JoinedEvents = joinedEvents

	return eventFormatter

}
func FormatEventDetail(event entity.Event) EventDetailFormatter {
	eventDetailFormatter := EventDetailFormatter{
		ID:            event.ID,
		Label:         event.Label,
		Name:          event.Name,
		Date:          event.Date,
		StartTime:     event.StartTime,
		EndTime:       event.EndTime,
		MeetURL:       event.MeetURL,
		Platform:      event.Platform,
		About:         event.About,
		Description:   event.Description,
		Documentation: event.Documentation,
		Status:        event.Status,
		Mentor: MentorFormatter{
			Name:       event.Mentor.Name,
			Occupation: event.Mentor.Occupation,
			AvatarURL:  event.Mentor.AvatarURL,
		},
	}

	return eventDetailFormatter

}
func FormatCalendarEvents(events []entity.CalendarEvent) []CalendarEventFormatter {
	calendarEventsFormatter := []CalendarEventFormatter{}

	for _, event := range events {
		calendarEventFormatter := FormatCalendarEvent(event)
		calendarEventsFormatter = append(calendarEventsFormatter, calendarEventFormatter)
	}

	return calendarEventsFormatter
}

func FormatEvents(events []entity.Event) []EventFormatter {
	eventsFormatter := []EventFormatter{}

	for _, event := range events {
		eventFormatter := FormatEvent(event)
		eventsFormatter = append(eventsFormatter, eventFormatter)
	}

	return eventsFormatter
}
