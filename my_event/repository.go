package myevent

import (
	"api-dunia-coding/event"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(memberID int) ([]event.Event, error)
	FindByStatus(memberID int) ([]event.Event, error)
	Update(eventID int, memberID int, presence bool) error
	CheckIsPresenced(eventID int, memberID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(memberID int) ([]event.Event, error) {
	var events []event.Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Table("events").Joins("JOIN joined_events ON events.id = joined_events.event_id").Where("joined_events.member_id = ?", memberID).Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}
func (r *repository) FindByStatus(memberID int) ([]event.Event, error) {
	var events []event.Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Table("events").Joins("JOIN joined_events ON events.id = joined_events.event_id").Where("joined_events.member_id = ?", memberID).Where("status = ?", "Upcoming").Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) Update(eventID int, memberID int, presence bool) error {

	var joinedEvent event.JoinedEvents

	result := r.db.Where("event_id = ? AND member_id = ?", eventID, memberID).First(&joinedEvent)

	if result.Error != nil {
		return result.Error
	}

	// Update nilai kolom Presence
	joinedEvent.Presence = presence

	// Simpan perubahan ke dalam database
	result = r.db.Save(&joinedEvent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) CheckIsPresenced(eventID int, memberID int) (bool, error) {
	var count int64
	err := r.db.Model(&event.JoinedEvents{}).Where("event_id = ? AND member_id = ? AND presence = ?", eventID, memberID, true).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
