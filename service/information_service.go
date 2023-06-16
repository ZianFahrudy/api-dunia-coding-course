package service

import (
	"api-dunia-coding/domain/entity"
	"api-dunia-coding/domain/repository"
	"context"
)

type InformationService interface {
	GetAll(ctx context.Context) ([]entity.Information, error)
	GetByID(ctx context.Context, ID int) (entity.Information, error)
	Delete(ctx context.Context, ID int) (bool, error)
	DeleteAll(ctx context.Context, ID []int) (bool, error)
}

type informationServiceImpl struct {
	repository repository.InformationRepository
}

func NewInformationServiceImpl(repository repository.InformationRepository) InformationService {
	return &informationServiceImpl{repository}
}

func (s *informationServiceImpl) GetAll(ctx context.Context) ([]entity.Information, error) {
	informations, err := s.repository.FindAll(ctx)

	if err != nil {
		return informations, err
	}

	return informations, nil
}

func (s *informationServiceImpl) GetByID(ctx context.Context, ID int) (entity.Information, error) {
	information, err := s.repository.FindByID(ctx, ID)

	if err != nil {
		return information, err
	}

	return information, nil
}

func (s *informationServiceImpl) Delete(ctx context.Context, ID int) (bool, error) {
	isDeleted, err := s.repository.Delete(ctx, ID)

	if err != nil {
		return false, err
	}

	return isDeleted, nil
}

func (s *informationServiceImpl) DeleteAll(ctx context.Context, ID []int) (bool, error) {
	isDeleted, err := s.repository.DeleteAll(ctx, ID)

	if err != nil {
		return false, err
	}

	return isDeleted, nil
}
