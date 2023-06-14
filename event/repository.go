package event

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindByID(ID int) (Event, error)
	Update(event Event) (Event, error)
	UpdateStatusEvent(ID int, status string) error
	FindByDate() ([]Event, error)
	FindByStatus(statusEvent string) ([]Event, error)
	FindByGroupDate() ([]CalendarEvent, error)
	SaveJoinEvent(joinedEvents JoinedEvents) (JoinedEvents, error)
	CheckEventMember(eventID int, memberID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Find(&events).Error

	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) FindByID(ID int) (Event, error) {
	var event Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Where("id = ?", ID).Find(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) Update(event Event) (Event, error) {
	err := r.db.Save(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) UpdateStatusEvent(ID int, status string) error {

	err := r.db.Model(&Event{}).Where("id = ?", ID).Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByDate() ([]Event, error) {
	var events []Event

	startOfWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	startWeek := startOfWeek.Format("2006-01-02")
	endWeek := endOfWeek.Format("2006-01-02")
	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Where("date BETWEEN ? AND ?", startWeek, endWeek).Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) FindByStatus(statusEvent string) ([]Event, error) {
	var events []Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Where("status = ?", statusEvent).Find(&events).Error

	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) FindByGroupDate() ([]CalendarEvent, error) {
	var events []Event
	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").
		Table("events").Select("date, id, name, label, mentor_id, start_time, end_time, status").Order("date desc").Find(&events).Error
	if err != nil {
		panic(err)
	}
	calendarEvents := make(map[string][]Event)

	for _, event := range events {
		if _, ok := calendarEvents[event.Date]; !ok {
			calendarEvents[event.Date] = []Event{}
		}
		calendarEvents[event.Date] = append(calendarEvents[event.Date], event)
	}
	var resultEvents []CalendarEvent
	for date, events := range calendarEvents {
		resultEvents = append(resultEvents, CalendarEvent{Date: date, Event: events})
	}
	return resultEvents, nil
}

func (r *repository) SaveJoinEvent(joinedEvent JoinedEvents) (JoinedEvents, error) {

	err := r.db.Table("joined_events").Create(&joinedEvent).Error
	if err != nil {
		return joinedEvent, err
	}

	return joinedEvent, nil
}

func (r *repository) CheckEventMember(eventID int, memberID int) (bool, error) {
	var count int64
	err := r.db.Table("joined_events").Where("event_id = ? AND member_id = ?", eventID, memberID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
