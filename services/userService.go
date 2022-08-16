package services

import (
	"log"
	"mitramas_test/models"
	"mitramas_test/repositories"
)

type UserService interface {
	ReadUser() (*[]models.UserResponse, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) ReadUser() (*[]models.UserResponse, error) {
	users, err := s.userRepository.ReadUser()
	if err != nil {
		log.Printf("Error get data users with err: %s", err)
		return nil, err
	}

	return users, nil
}
