package repository

import (
	"database/sql"
	"fmt"
	"jukeBE/internal/model"
	"strings"
)

type EmployeeRepository interface {
	GetAll(params model.PaginationQuery) ([]*model.Employee, int, error)
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

func (r *postgresEmployeeRepository) GetAll(params model.PaginationQuery) ([]*model.Employee, int, error) {
	var conditions []string
	var args []interface{}
	argID := 1

	// Filter by Position
	if params.Position != "" {
		conditions = append(conditions, fmt.Sprintf("position = $%d", argID))
		args = append(args, params.Position)
		argID++
	}

	// Search
	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR email ILIKE $%d)", argID, argID+1))
		args = append(args, searchPattern, searchPattern)
		argID += 2
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// Count Total
	countQuery := "SELECT COUNT(*) FROM employees" + whereClause
	var total int
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Sorting
	sortMap := map[string]string{
		"name":       "name",
		"salary":     "salary",
		"created_at": "created_at",
		"id":         "id",
	}
	sortBy, ok := sortMap[strings.ToLower(params.SortBy)]
	if !ok {
		sortBy = "created_at"
	}
	sortOrder := "DESC"
	if strings.ToUpper(params.SortOrder) == "ASC" {
		sortOrder = "ASC"
	}

	// Pagination
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (params.Page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	query := fmt.Sprintf(`SELECT id, name, email, position, salary, profile_picture, created_at FROM employees%s ORDER BY %s %s LIMIT $%d OFFSET $%d`, whereClause, sortBy, sortOrder, argID, argID+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		e := &model.Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.Salary, &e.ProfilePicture, &e.CreatedAt); err != nil {
			return nil, 0, err
		}
		employees = append(employees, e)
	}
	return employees, total, nil
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
