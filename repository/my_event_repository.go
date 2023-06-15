package repository

import (
	"api-dunia-coding/entity"
	"context"
)

type MyEventRepository interface {
	FindAll(ctx context.Context, memberID int) ([]entity.Event, error)
	FindByStatus(ctx context.Context, memberID int) ([]entity.Event, error)
	Update(ctx context.Context, eventID int, memberID int, presence bool) error
	CheckIsPresenced(ctx context.Context, eventID int, memberID int) (bool, error)
}
