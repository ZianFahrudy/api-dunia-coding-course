package repository

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/exception"
	"context"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *authRepositoryImpl {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) Save(ctx context.Context, member entity.Member) entity.Member {

	err := r.db.WithContext(ctx).Create(&member).Error
	exception.PanicIfNeeded(err)

	return member
}

func (r *authRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.Member, error) {
	var member entity.Member

	err := r.db.WithContext(ctx).Where("email = ?", email).Find(&member).Error
	exception.PanicIfNeeded(err)

	return member, nil
}

func (r *authRepositoryImpl) FindEmailExist(ctx context.Context, email string) bool {
	var member entity.Member

	err := r.db.WithContext(ctx).First(&member, "email = ?", email).Error
	if err != nil {
		return false
	}

	return true
}

func (r *authRepositoryImpl) FindByID(ctx context.Context, ID int) (entity.Member, error) {
	var member entity.Member
	err := r.db.WithContext(ctx).Where("id = ?", ID).Find(&member).Error
	exception.PanicIfNeeded(err)

	return member, nil
}

func (r *authRepositoryImpl) Update(ctx context.Context, member entity.Member) (entity.Member, error) {
	err := r.db.WithContext(ctx).Save(&member).Error
	exception.PanicIfNeeded(err)

	return member, nil
}
