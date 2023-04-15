package event

type EventFormatter struct {
	ID        int    `json:"id"`
	Label     string `json:"label"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	MeetURL   string `json:"meet_url"`
	// Platform      string                 `json:"platform"`
	// About         string                 `json:"about"`
	// Description   string                 `json:"description"`
	// Documentation string                 `json:"documentation"`
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

func FormatCalendarEvent(event CalendarEvent) CalendarEventFormatter {
	calendarEventFormatter := CalendarEventFormatter{}
	calendarEventFormatter.Date = event.Date
	eventsCalendar := []EventCalendarFormatter{}

	for _, event := range event.Event {
		eventCalendarFormatter := EventCalendarFormatter{}

		eventCalendarFormatter.ID = event.ID
		eventCalendarFormatter.Date = event.Date
		eventCalendarFormatter.Label = event.Label
		eventCalendarFormatter.Name = event.Name
		eventCalendarFormatter.StartTime = event.StartTime
		eventCalendarFormatter.EndTime = event.EndTime
		eventCalendarFormatter.Status = event.Status

		mentorFormatter := MentorFormatter{}

		mentorFormatter.Name = event.Mentor.Name
		mentorFormatter.Occupation = event.Mentor.Occupation
		mentorFormatter.AvatarURL = event.Mentor.AvatarURL

		eventCalendarFormatter.Mentor = mentorFormatter

		calendarEventFormatter.Events = eventsCalendar
		eventsCalendar = append(eventsCalendar, eventCalendarFormatter)
	}
	calendarEventFormatter.Events = eventsCalendar

	return calendarEventFormatter
}

func FormatEvent(event Event) EventFormatter {
	eventFormatter := EventFormatter{}
	eventFormatter.ID = event.ID
	eventFormatter.Label = event.Label
	eventFormatter.Name = event.Name
	eventFormatter.Date = event.Date
	eventFormatter.StartTime = event.StartTime
	eventFormatter.EndTime = event.EndTime
	eventFormatter.MeetURL = event.MeetURL
	// eventFormatter.Platform = event.Platform
	// eventFormatter.About = event.Platform
	// eventFormatter.Description = event.Description
	// eventFormatter.Documentation = event.Documentation

	eventFormatter.Status = event.Status

	mentorFormatter := MentorFormatter{}

	mentorFormatter.Name = event.Mentor.Name
	mentorFormatter.Occupation = event.Mentor.Occupation
	mentorFormatter.AvatarURL = event.Mentor.AvatarURL

	eventFormatter.Mentor = mentorFormatter

	joinedEvents := []JoinedEventFormatter{}

	for _, joinedEvent := range event.JoinedEvents {
		memberEventFormatter := MemberEventFormatter{}
		joinedEventFormatter := JoinedEventFormatter{}

		// joinedEventFormatter.MemberID = joinedEvent.MemberID

		memberEventFormatter.Email = joinedEvent.Member.Email
		memberEventFormatter.Name = joinedEvent.Member.Name

		joinedEventFormatter.Member = memberEventFormatter

		joinedEvents = append(joinedEvents, joinedEventFormatter)
	}

	eventFormatter.JoinedEvents = joinedEvents

	return eventFormatter

}
func FormatEventDetail(event Event) EventDetailFormatter {
	eventFormatter := EventDetailFormatter{}
	eventFormatter.ID = event.ID
	eventFormatter.Label = event.Label
	eventFormatter.Name = event.Name
	eventFormatter.Date = event.Date
	eventFormatter.StartTime = event.StartTime
	eventFormatter.EndTime = event.EndTime
	eventFormatter.MeetURL = event.MeetURL
	eventFormatter.Platform = event.Platform
	eventFormatter.About = event.Platform
	eventFormatter.Description = event.Description
	eventFormatter.Documentation = event.Documentation

	eventFormatter.Status = event.Status

	mentorFormatter := MentorFormatter{}

	mentorFormatter.Name = event.Mentor.Name
	mentorFormatter.Occupation = event.Mentor.Occupation
	mentorFormatter.AvatarURL = event.Mentor.AvatarURL

	eventFormatter.Mentor = mentorFormatter

	joinedEvents := []JoinedEventFormatter{}

	for _, joinedEvent := range event.JoinedEvents {
		memberEventFormatter := MemberEventFormatter{}
		joinedEventFormatter := JoinedEventFormatter{}

		// joinedEventFormatter.MemberID = joinedEvent.MemberID

		memberEventFormatter.Email = joinedEvent.Member.Email
		memberEventFormatter.Name = joinedEvent.Member.Name

		joinedEventFormatter.Member = memberEventFormatter

		joinedEvents = append(joinedEvents, joinedEventFormatter)
	}

	// eventFormatter.JoinedEvent = joinedEventTime

	return eventFormatter

}
func FormatCalendarEvents(events []CalendarEvent) []CalendarEventFormatter {
	calendarEventsFormatter := []CalendarEventFormatter{}

	for _, event := range events {
		calendarEventFormatter := FormatCalendarEvent(event)
		calendarEventsFormatter = append(calendarEventsFormatter, calendarEventFormatter)
	}

	return calendarEventsFormatter
}

func FormatEvents(events []Event) []EventFormatter {
	eventsFormatter := []EventFormatter{}

	for _, event := range events {
		eventFormatter := FormatEvent(event)
		eventsFormatter = append(eventsFormatter, eventFormatter)
	}

	return eventsFormatter
}