package myevent

import "api-dunia-coding/event"

type Service interface {
	GetMyEvents(memberID int) ([]event.Event, error)
	GetUpcomingMyEvents(memberID int) ([]event.Event, error)
	Presence(input PresenceInput, memberID int) error
	CheckIsPresenced(input PresenceInput, memberID int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMyEvents(memberID int) ([]event.Event, error) {
	events, err := s.repository.FindAll(memberID)

	if err != nil {
		return events, err
	}

	return events, nil
}
func (s *service) GetUpcomingMyEvents(memberID int) ([]event.Event, error) {
	events, err := s.repository.FindByStatus(memberID)

	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *service) Presence(input PresenceInput, memberID int) error {
	err := s.repository.Update(input.EventID, memberID, true)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CheckIsPresenced(input PresenceInput, memberID int) (bool, error) {
	_, err := s.repository.CheckIsPresenced(input.EventID, memberID)

	if err != nil {
		return false, err
	}
	return true, nil
}
