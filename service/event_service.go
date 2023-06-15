package service

import (
	"api-dunia-coding/data/model"
	"api-dunia-coding/domain/entity"
	"context"
)

type EventService interface {
	GetEvents(ctx context.Context) ([]entity.Event, error)
	UpdateStatusEvent(ctx context.Context, ID int, status string) error
	GetEventByID(ctx context.Context, input model.GetEventDetailInput) (entity.Event, error)
	GetEventOfWeek(ctx context.Context) ([]entity.Event, error)
	GetEventByStatus(ctx context.Context, status string) ([]entity.Event, error)
	GetCalendarEvents(ctx context.Context) ([]entity.CalendarEvent, error)
	JoinToEvent(ctx context.Context, input model.JoinEventInput) (entity.JoinedEvents, error)
	CheckEventMember(ctx context.Context, eventID int, memberID int) (bool, error)
}
