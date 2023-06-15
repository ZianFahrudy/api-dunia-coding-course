package service

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/model"
	"api-dunia-coding/repository"
	"context"
)

func NewEventServiceImpl(eventRepository repository.EventRepository) EventService {
	return &eventServiceImpl{EventRepository: eventRepository}
}

type eventServiceImpl struct {
	EventRepository repository.EventRepository
}

func (s *eventServiceImpl) GetEvents(ctx context.Context) ([]entity.Event, error) {

	events, err := s.EventRepository.FindAll(ctx)
	if err != nil {
		return events, err
	}

	return events, nil

}

func (s *eventServiceImpl) UpdateStatusEvent(ctx context.Context, ID int, status string) error {
	err := s.EventRepository.UpdateStatusEvent(ctx, ID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *eventServiceImpl) GetEventByID(ctx context.Context, input model.GetEventDetailInput) (entity.Event, error) {
	event, err := s.EventRepository.FindByID(ctx, input.ID)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (s *eventServiceImpl) GetEventOfWeek(ctx context.Context) ([]entity.Event, error) {
	events, err := s.EventRepository.FindByDate(ctx)
	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *eventServiceImpl) GetEventByStatus(ctx context.Context, status string) ([]entity.Event, error) {
	events, err := s.EventRepository.FindByStatus(ctx, status)
	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *eventServiceImpl) GetCalendarEvents(ctx context.Context) ([]entity.CalendarEvent, error) {
	calendarEvents, err := s.EventRepository.FindByGroupDate(ctx)
	if err != nil {
		return calendarEvents, err
	}

	return calendarEvents, nil
}

func (s *eventServiceImpl) JoinToEvent(ctx context.Context, input model.JoinEventInput) (entity.JoinedEvents, error) {
	joinedEvent := entity.JoinedEvents{}

	joinedEvent.EventID = input.EventID
	joinedEvent.MemberID = input.MemberID

	joinedEvent, err := s.EventRepository.SaveJoinEvent(ctx, joinedEvent)
	if err != nil {
		return joinedEvent, err
	}

	return joinedEvent, nil
}

func (s *eventServiceImpl) CheckEventMember(ctx context.Context, eventID int, memberID int) (bool, error) {
	isPresenced, err := s.EventRepository.CheckEventMember(ctx, eventID, memberID)
	if err != nil {
		return false, err
	}

	return isPresenced, nil
}
