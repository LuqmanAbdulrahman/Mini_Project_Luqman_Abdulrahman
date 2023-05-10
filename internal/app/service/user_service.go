package service

import (
	"errors"

	"your-project-name/internal/app/model"
	"your-project-name/internal/app/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(userReq *model.UserRequest) (*model.User, error) {
	// Check if the user already exists
	existingUser, err := s.userRepo.FindByEmail(userReq.Email)
	if err != nil && !errors.Is(err, repository.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Create new user
	user := model.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	err = s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUser(userID string) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(userID string, userReq *model.UserRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Name = userReq.Name
	user.Email = userReq.Email
	user.Password = userReq.Password

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(userID string) error {
	err := s.userRepo.DeleteByID(userID)
	if err != nil {
		return err
	}
	return nil
}

//membuat struktur UserService yang memiliki metode-metode untuk melakukan operasi CRUD pada entitas User.
//juga menambahkan validasi agar tidak ada duplikasi data ketika membuat data baru.
