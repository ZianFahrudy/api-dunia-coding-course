package member

import "gorm.io/gorm"

type Repository interface {
	Save(member Member) (Member, error)
	FindByEmail(email string) (Member, error)
	FindByID(ID int) (Member, error)
	Update(member Member) (Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(member Member) (Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repository) FindByEmail(email string) (Member, error) {
	var member Member
	err := r.db.Where("email = ?", email).Find(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repository) FindByID(ID int) (Member, error) {
	var member Member
	err := r.db.Where("id = ?", ID).Find(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repository) Update(member Member) (Member, error) {
	err := r.db.Save(&member).Error

	if err != nil {
		return member, err
	}

	return member, nil
}
