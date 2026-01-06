package service

import (
	"errors"
	"jukeBE/internal/model"
	"jukeBE/internal/repository"
)

type AuthService interface {
	Login(email, password string) (*model.Admin, error)
	ValidateCredentials(username, password string) bool
}

type authService struct {
	adminRepo repository.AdminRepository
}

func NewAuthService(adminRepo repository.AdminRepository) AuthService {
	return &authService{adminRepo: adminRepo}
}

func (s *authService) Login(email, password string) (*model.Admin, error) {
	admin, err := s.adminRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// In a real app, verify hash. checking plain text for this simple auth task
	if admin.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return admin, nil
}

func (s *authService) ValidateCredentials(username, password string) bool {
	admin, err := s.adminRepo.GetByEmail(username)
	if err != nil {
		return false
	}

	// In a real app, verify hash. checking plain text for this simple auth task
	return admin.Password == password
}
