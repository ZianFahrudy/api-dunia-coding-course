package repository

import (
	"api-dunia-coding/entity"
	"context"
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
