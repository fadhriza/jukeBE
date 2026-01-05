package repository

import (
	"database/sql"
	"jukeBE/internal/model"
)

type AdminRepository interface {
	GetByEmail(email string) (*model.Admin, error)
}

type postgresAdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &postgresAdminRepository{DB: db}
}

func (r *postgresAdminRepository) GetByEmail(email string) (*model.Admin, error) {
	query := `SELECT id, name, email, password FROM admins WHERE email = $1`
	admin := &model.Admin{}
	err := r.DB.QueryRow(query, email).Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Password)
	if err != nil {
		return nil, err
	}
	return admin, nil
}
