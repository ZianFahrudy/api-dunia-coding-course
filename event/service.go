package event

type Service interface {
	GetEvents() ([]Event, error)
	UpdateStatusEvent(ID int, status string) error
	GetEventByID(input GetEventDetailInput) (Event, error)
	GetEventOfWeek() ([]Event, error)
	GetEventByStatus(status string) ([]Event, error)
	GetCalendarEvents() ([]CalendarEvent, error)
	JoinToEvent(input JoinEventInput) (JoinedEvents, error)
	CheckEventMember(eventID int, memberID int) (bool, error)
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

func (s *service) GetEventOfWeek() ([]Event, error) {
	events, err := s.repository.FindByDate()
	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *service) GetEventByStatus(status string) ([]Event, error) {
	events, err := s.repository.FindByStatus(status)
	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *service) GetCalendarEvents() ([]CalendarEvent, error) {
	calendarEvents, err := s.repository.FindByGroupDate()
	if err != nil {
		return calendarEvents, err
	}

	return calendarEvents, nil
}

func (s *service) JoinToEvent(input JoinEventInput) (JoinedEvents, error) {
	joinedEvent := JoinedEvents{}

	joinedEvent.EventID = input.EventID
	joinedEvent.MemberID = input.MemberID

	joinedEvent, err := s.repository.SaveJoinEvent(joinedEvent)
	if err != nil {
		return joinedEvent, err
	}

	return joinedEvent, nil
}

func (s *service) CheckEventMember(eventID int, memberID int) (bool, error) {
	isPresenced, err := s.repository.CheckEventMember(eventID, memberID)
	if err != nil {
		return false, err
	}

	return isPresenced, nil
}
