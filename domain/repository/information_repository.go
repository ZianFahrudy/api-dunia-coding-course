package repository

import (
	"api-dunia-coding/domain/entity"
	"context"

	"gorm.io/gorm"
)

type InformationRepository interface {
	FindAll(ctx context.Context) ([]entity.Information, error)
	FindByID(ctx context.Context, ID int) (entity.Information, error)
	Delete(ctx context.Context, ID int) (bool, error)
	DeleteAll(ctx context.Context, ID []int) (bool, error)
}

type informationRepositoryImpl struct {
	db *gorm.DB
}

func NewInformationRepositoryImpl(db *gorm.DB) *informationRepositoryImpl {
	return &informationRepositoryImpl{db}
}

func (r *informationRepositoryImpl) FindAll(ctx context.Context) ([]entity.Information, error) {
	var informations []entity.Information

	err := r.db.WithContext(ctx).Table("informations").Find(&informations).Error
	if err != nil {
		return informations, err
	}

	return informations, nil
}

func (r *informationRepositoryImpl) FindByID(ctx context.Context, ID int) (entity.Information, error) {
	var information entity.Information

	err := r.db.WithContext(ctx).Table("informations").Where("id = ?", ID).Find(&information).Error
	if err != nil {
		return information, err
	}

	return information, nil
}

func (r *informationRepositoryImpl) Delete(ctx context.Context, ID int) (bool, error) {
	var information entity.Information

	err := r.db.WithContext(ctx).Table("informations").Where("id = ?", ID).Delete(&information).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *informationRepositoryImpl) DeleteAll(ctx context.Context, ID []int) (bool, error) {
	var information entity.Information

	err := r.db.WithContext(ctx).Table("informations").Where("id IN (?)", ID).Delete(&information).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
