package service

import (
	"errors"
	"jukeBE/internal/model"
	"jukeBE/internal/repository"
	"math"
)

type EmployeeService interface {
	GetAllEmployees(params model.PaginationQuery) (*model.PaginatedResponse, error)
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

func (s *employeeService) GetAllEmployees(params model.PaginationQuery) (*model.PaginatedResponse, error) {
	employees, total, err := s.repo.GetAll(params)
	if err != nil {
		return nil, err
	}

	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &model.PaginatedResponse{
		Data:       employees,
		Total:      total,
		Page:       params.Page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
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
