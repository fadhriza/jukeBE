package service

import (
	"errors"
	"jukeBE/internal/model"
	"jukeBE/internal/repository"
)

type EmployeeService interface {
	GetAllEmployees() ([]*model.Employee, error)
	GetEmployee(id int64) (*model.Employee, error)
	CreateEmployee(employee *model.Employee) error
	UpdateEmployee(id int64, employee *model.Employee) error
	DeleteEmployee(id int64) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllEmployees() ([]*model.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetEmployee(id int64) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) CreateEmployee(e *model.Employee) error {
	// Validation
	if e.Salary <= 0 {
		return errors.New("salary must be greater than 0")
	}
	exists, err := s.repo.EmailExists(e.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	return s.repo.Create(e)
}

func (s *employeeService) UpdateEmployee(id int64, e *model.Employee) error {
	// Check existence
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
    
    // Validate salary if being updated
    if e.Salary <= 0 {
        return errors.New("salary must be greater than 0")
    }

	// Update fields
	existing.Name = e.Name
	existing.Email = e.Email
	existing.Position = e.Position
	existing.Salary = e.Salary
    if e.ProfilePicture != "" {
    	existing.ProfilePicture = e.ProfilePicture
    }

	return s.repo.Update(existing)
}

func (s *employeeService) DeleteEmployee(id int64) error {
	return s.repo.Delete(id)
}
