package service

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/model"
	"context"
)

type AuthService interface {
	RegisterMember(ctx context.Context, body model.RegisterMemberBody) entity.Member
	Login(ctx context.Context, body model.LoginBody) (entity.Member, error)
	CheckEmailOrPasswordValid(ctx context.Context, body model.LoginBody) (bool, error)
	GetMemberByID(ctx context.Context, body int) (entity.Member, error)
	CheckEmailAvailable(ctx context.Context, input string) (bool, error)
	SaveAvatar(ctx context.Context, ID int, fileLocation string) (entity.Member, error)
}
