package service

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/model"
	"context"
)

type MyEventService interface {
	GetMyEvents(ctx context.Context, memberID int) ([]entity.Event, error)
	GetUpcomingMyEvents(ctx context.Context, memberID int) ([]entity.Event, error)
	Presence(ctx context.Context, input model.PresenceInput, memberID int) error
	CheckIsPresenced(ctx context.Context, input model.PresenceInput, memberID int) (bool, error)
}