package event

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Event, error)
	FindByID(ID int) (Event, error)
	Update(event Event) (Event, error)
	UpdateStatusEvent(ID int, status string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event

	err := r.db.Preload("Mentor").Preload("JoinedEvent.Member").Find(&events).Error

	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) FindByID(ID int) (Event, error) {
	var event Event

	err := r.db.Preload("Mentor").Preload("JoinedEvent.Member").Where("id = ?", ID).Find(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) Update(event Event) (Event, error) {
	err := r.db.Save(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) UpdateStatusEvent(ID int, status string) error {

	err := r.db.Model(&Event{}).Where("id = ?", ID).Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}
