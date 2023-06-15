package service

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/model"
	"api-dunia-coding/repository"
	"context"
)

func NewMyEventServiceImpl(myEventrepository repository.MyEventRepository) MyEventService {
	return &myEventServiceImpl{myEventrepository}
}

type myEventServiceImpl struct {
	repository repository.MyEventRepository
}

func (s *myEventServiceImpl) GetMyEvents(ctx context.Context, memberID int) ([]entity.Event, error) {
	events, err := s.repository.FindAll(ctx, memberID)

	if err != nil {
		return events, err
	}

	return events, nil
}
func (s *myEventServiceImpl) GetUpcomingMyEvents(ctx context.Context, memberID int) ([]entity.Event, error) {
	events, err := s.repository.FindByStatus(ctx, memberID)

	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *myEventServiceImpl) Presence(ctx context.Context, input model.PresenceInput, memberID int) error {
	err := s.repository.Update(ctx, input.EventID, memberID, true)

	if err != nil {
		return err
	}

	return nil
}

func (s *myEventServiceImpl) CheckIsPresenced(ctx context.Context, input model.PresenceInput, memberID int) (bool, error) {
	_, err := s.repository.CheckIsPresenced(ctx, input.EventID, memberID)

	if err != nil {
		return false, err
	}
	return true, nil
}