package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterMember(input RegisterMemberInput) (Member, error)
	Login(input LoginInput) (Member, error)
	CheckEmailOrPasswordValid(input LoginInput) (bool, error)
	GetMemberByID(ID int) (Member, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (Member, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterMember(input RegisterMemberInput) (Member, error) {
	member := Member{}

	member.Name = input.Name
	member.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return member, err
	}

	member.PasswordHash = string(passwordHash)

	newMember, err := s.repository.Save(member)
	if err != nil {
		return newMember, err
	}

	return newMember, nil

}

func (s *service) Login(input LoginInput) (Member, error) {

	member, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, errors.New("Tidak ada member yang menggunakan email ini")
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(input.Password))
	if err != nil {
		return member, err
	}

	return member, nil

}

func (s *service) CheckEmailOrPasswordValid(input LoginInput) (bool, error) {
	member, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if member.ID == 0 {
		return false, errors.New("Tidak ada member yang menggunakan email ini")
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(input.Password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GetMemberByID(ID int) (Member, error) {
	member, err := s.repository.FindByID(ID)
	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, errors.New("No member found on that ID")

	}

	return member, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {

	member, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return false, err
	}

	if member.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (Member, error) {
	member, err := s.repository.FindByID(ID)

	if err != nil {
		return member, err
	}

	member.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(member)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
