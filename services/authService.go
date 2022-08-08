package services

import (
	"log"
	"mitramas_test/models"
	"mitramas_test/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.UserRegister) error
	Login(user *models.UserLogin) (*models.UserLogin, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) *authService {
	return &authService{repository}
}

func (s *authService) Register(user *models.UserRegister) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		log.Printf("Error hashing password with err: %s", err)
		return err
	}
	user.Password = string(password)

	userRegister := models.UserRegister{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.authRepository.Register(&userRegister)
	if err != nil {
		if err != nil {
			log.Printf("Error register user to database with err: %s", err)
			return err
		}
	}

	return nil
}

func (s *authService) Login(user *models.UserLogin) (*models.UserLogin, error) {
	userLogin := models.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	}
	userData, err := s.authRepository.Login(&userLogin)
	if err != nil {
		log.Printf("Error login to database with err: %s", err)
		return nil, err
	}

	return userData, nil
}
