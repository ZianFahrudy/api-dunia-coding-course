package service

import (
	"api-dunia-coding/entity"
	"api-dunia-coding/model"
	"api-dunia-coding/repository"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthServiceImpl(authRepository repository.AuthRepository) AuthService {
	return &authServiceImpl{AuthRepository: authRepository}
}

type authServiceImpl struct {
	AuthRepository repository.AuthRepository
}

func (s *authServiceImpl) RegisterMember(ctx context.Context, body model.RegisterMemberBody) entity.Member {
	member := entity.Member{}

	member.Name = body.Name
	member.Email = body.Email

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)

	member.PasswordHash = string(passwordHash)

	newMember := s.AuthRepository.Save(ctx, member)

	return newMember

}

func (s *authServiceImpl) Login(ctx context.Context, body model.LoginBody) (entity.Member, error) {

	member, err := s.AuthRepository.FindByEmail(ctx, body.Email)
	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, err
	}

	bcrypt.CompareHashAndPassword([]byte(member.PasswordHash), []byte(body.Password))

	return member, nil

}

func (s *authServiceImpl) CheckEmailOrPasswordValid(ctx context.Context, input model.LoginBody) (bool, error) {
	member, err := s.AuthRepository.FindByEmail(ctx, input.Email)
	fmt.Println(member.PasswordHash)
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

func (s *authServiceImpl) GetMemberByID(ctx context.Context, body int) (entity.Member, error) {
	member, err := s.AuthRepository.FindByID(ctx, body)
	if err != nil {
		return member, err
	}

	if member.ID == 0 {
		return member, errors.New("No member found on that ID")

	}

	return member, nil
}

func (s *authServiceImpl) SaveAvatar(ctx context.Context, ID int, fileLocation string) (entity.Member, error) {
	member, err := s.AuthRepository.FindByID(ctx, ID)
	fmt.Println(member.Email)

	if err != nil {
		return member, err
	}

	member.AvatarFileName = fileLocation

	updatedUser, errs := s.AuthRepository.Update(ctx, member)
	if errs != nil {
		return updatedUser, errs
	}

	return updatedUser, nil
}

func (s *authServiceImpl) CheckEmailAvailable(ctx context.Context, input string) (bool, error) {

	isExist := s.AuthRepository.FindEmailExist(ctx, input)

	return isExist, nil
}
