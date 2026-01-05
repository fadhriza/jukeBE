package repository

import (
	"database/sql"
	"jukeBE/internal/model"
)

type EmployeeRepository interface {
	GetAll() ([]*model.Employee, error)
	GetByID(id int64) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(employee *model.Employee) error
	Delete(id int64) error
	EmailExists(email string) (bool, error)
}

type postgresEmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &postgresEmployeeRepository{DB: db}
}

func (r *postgresEmployeeRepository) GetAll() ([]*model.Employee, error) {
	query := `SELECT id, name, email, position, salary, profile_picture, created_at FROM employees`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		e := &model.Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.Salary, &e.ProfilePicture, &e.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func (r *postgresEmployeeRepository) GetByID(id int64) (*model.Employee, error) {
	query := `SELECT id, name, email, position, salary, profile_picture, created_at FROM employees WHERE id = $1`
	e := &model.Employee{}
	err := r.DB.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.Salary, &e.ProfilePicture, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *postgresEmployeeRepository) Create(e *model.Employee) error {
	query := `INSERT INTO employees (name, email, position, salary, profile_picture) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	return r.DB.QueryRow(query, e.Name, e.Email, e.Position, e.Salary, e.ProfilePicture).Scan(&e.ID, &e.CreatedAt)
}

func (r *postgresEmployeeRepository) Update(e *model.Employee) error {
	query := `UPDATE employees SET name=$1, email=$2, position=$3, salary=$4, profile_picture=$5 WHERE id=$6`
	_, err := r.DB.Exec(query, e.Name, e.Email, e.Position, e.Salary, e.ProfilePicture, e.ID)
	return err
}

func (r *postgresEmployeeRepository) Delete(id int64) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *postgresEmployeeRepository) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM employees WHERE email=$1)`
	err := r.DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}
