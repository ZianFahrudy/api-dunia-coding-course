package repository

import (
	"api-dunia-coding/domain/entity"
	"context"
	"time"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll(ctx context.Context) ([]entity.Event, error)
	FindByID(ctx context.Context, ID int) (entity.Event, error)
	Update(ctx context.Context, event entity.Event) (entity.Event, error)
	UpdateStatusEvent(ctx context.Context, ID int, status string) error
	FindByDate(ctx context.Context) ([]entity.Event, error)
	FindByStatus(ctx context.Context, statusEvent string) ([]entity.Event, error)
	FindByGroupDate(ctx context.Context) ([]entity.CalendarEvent, error)
	SaveJoinEvent(ctx context.Context, joinedEvents entity.JoinedEvents) (entity.JoinedEvents, error)
	CheckEventMember(ctx context.Context, eventID int, memberID int) (bool, error)
}

type eventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepositoryImpl(db *gorm.DB) *eventRepositoryImpl {
	return &eventRepositoryImpl{db}
}

func (r *eventRepositoryImpl) FindAll(ctx context.Context) ([]entity.Event, error) {
	var events []entity.Event

	err := r.db.WithContext(ctx).Preload("Mentor").Preload("JoinedEvents.Member").Find(&events).Error

	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *eventRepositoryImpl) FindByID(ctx context.Context, ID int) (entity.Event, error) {
	var event entity.Event

	err := r.db.WithContext(ctx).Preload("Mentor").Preload("JoinedEvents.Member").Where("id = ?", ID).Find(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *eventRepositoryImpl) Update(ctx context.Context, event entity.Event) (entity.Event, error) {
	err := r.db.WithContext(ctx).Save(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *eventRepositoryImpl) UpdateStatusEvent(ctx context.Context, ID int, status string) error {

	err := r.db.WithContext(ctx).Model(&entity.Event{}).Where("id = ?", ID).Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepositoryImpl) FindByDate(ctx context.Context) ([]entity.Event, error) {
	var events []entity.Event

	startOfWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	startWeek := startOfWeek.Format("2006-01-02")
	endWeek := endOfWeek.Format("2006-01-02")
	err := r.db.WithContext(ctx).Preload("Mentor").Preload("JoinedEvents.Member").Where("date BETWEEN ? AND ?", startWeek, endWeek).Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *eventRepositoryImpl) FindByStatus(ctx context.Context, statusEvent string) ([]entity.Event, error) {
	var events []entity.Event

	err := r.db.WithContext(ctx).Preload("Mentor").Preload("JoinedEvents.Member").Where("status = ?", statusEvent).Find(&events).Error

	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *eventRepositoryImpl) FindByGroupDate(ctx context.Context) ([]entity.CalendarEvent, error) {
	var events []entity.Event
	err := r.db.WithContext(ctx).Preload("Mentor").Preload("JoinedEvents.Member").
		Table("events").Select("date, id, name, label, mentor_id, start_time, end_time, status").Order("date desc").Find(&events).Error
	if err != nil {
		panic(err)
	}
	calendarEvents := make(map[string][]entity.Event)

	for _, event := range events {
		if _, ok := calendarEvents[event.Date]; !ok {
			calendarEvents[event.Date] = []entity.Event{}
		}
		calendarEvents[event.Date] = append(calendarEvents[event.Date], event)
	}
	var resultEvents []entity.CalendarEvent
	for date, events := range calendarEvents {
		resultEvents = append(resultEvents, entity.CalendarEvent{Date: date, Event: events})
	}
	return resultEvents, nil
}

func (r *eventRepositoryImpl) SaveJoinEvent(ctx context.Context, joinedEvent entity.JoinedEvents) (entity.JoinedEvents, error) {

	err := r.db.WithContext(ctx).Table("joined_events").Create(&joinedEvent).Error
	if err != nil {
		return joinedEvent, err
	}

	return joinedEvent, nil
}

func (r *eventRepositoryImpl) CheckEventMember(ctx context.Context, eventID int, memberID int) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("joined_events").Where("event_id = ? AND member_id = ?", eventID, memberID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
