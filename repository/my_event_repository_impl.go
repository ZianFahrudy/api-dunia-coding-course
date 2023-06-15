package repository

import (
	"api-dunia-coding/entity"
	"context"

	"gorm.io/gorm"
)

type myEventRepositoryImpl struct {
	db *gorm.DB
}

func NewMyEventRepositoryImpl(db *gorm.DB) *myEventRepositoryImpl {
	return &myEventRepositoryImpl{db}
}

func (r *myEventRepositoryImpl) FindAll(ctx context.Context, memberID int) ([]entity.Event, error) {
	var events []entity.Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Table("events").Joins("JOIN joined_events ON events.id = joined_events.event_id").Where("joined_events.member_id = ?", memberID).Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}
func (r *myEventRepositoryImpl) FindByStatus(ctx context.Context, memberID int) ([]entity.Event, error) {
	var events []entity.Event

	err := r.db.Preload("Mentor").Preload("JoinedEvents.Member").Table("events").Joins("JOIN joined_events ON events.id = joined_events.event_id").Where("joined_events.member_id = ?", memberID).Where("status = ?", "Upcoming").Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *myEventRepositoryImpl) Update(ctx context.Context, eventID int, memberID int, presence bool) error {

	var joinedEvent entity.JoinedEvents

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

func (r *myEventRepositoryImpl) CheckIsPresenced(ctx context.Context, eventID int, memberID int) (bool, error) {
	var count int64
	err := r.db.Model(&entity.JoinedEvents{}).Where("event_id = ? AND member_id = ? AND presence = ?", eventID, memberID, true).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
