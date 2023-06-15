package repository

import (
	"api-dunia-coding/entity"
	"context"
)

type AuthRepository interface {
	Save(ctx context.Context, member entity.Member) entity.Member
	FindByEmail(ctx context.Context, email string) (entity.Member, error)
	FindEmailExist(ctx context.Context, email string) bool
	FindByID(ctx context.Context, ID int) (entity.Member, error)
	Update(ctx context.Context, member entity.Member) (entity.Member, error)
}
