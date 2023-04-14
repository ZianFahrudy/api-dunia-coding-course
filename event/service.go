package event

type Service interface {
	GetEvents() ([]Event, error)
	UpdateStatusEvent(ID int, status string) error
	GetEventByID(input GetEventDetailInput) (Event, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) GetEvents() ([]Event, error) {

	events, err := s.repository.FindAll()
	if err != nil {
		return events, err
	}

	return events, nil

}

func (s *service) UpdateStatusEvent(ID int, status string) error {
	err := s.repository.UpdateStatusEvent(ID, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetEventByID(input GetEventDetailInput) (Event, error) {
	event, err := s.repository.FindByID(input.ID)
	if err != nil {
		return event, err
	}
	return event, nil
}
