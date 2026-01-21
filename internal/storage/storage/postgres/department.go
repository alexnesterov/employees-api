package postgres

import (
	"context"
	"errors"

	"github.com/alexnesterov/employees-api/internal/models"
	"github.com/jackc/pgx/v5"
)

type DepartmentRepository struct {
	db *pgx.Conn
}

func NewDepartmentRepository(db *pgx.Conn) models.DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (r *DepartmentRepository) Create(dpt *models.Department) error {
	query := `INSERT INTO departments (code, name) VALUES ($1, $2) RETURNING code`
	params := []any{dpt.Code, dpt.Name}

	row := r.db.QueryRow(context.Background(), query, params...)
	err := row.Scan(&dpt.Code)
	if err != nil {
		return err
	}

	return nil
}

func (r *DepartmentRepository) List() ([]*models.Department, error) {
	query := `SELECT code, name FROM departments`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	departments, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*models.Department, error) {
		department := &models.Department{}
		err := row.Scan(&department.Code, &department.Name)
		if err != nil {
			return nil, err
		}
		return department, nil
	})

	return departments, err
}

func (r *DepartmentRepository) Read(code string) (*models.Department, error) {
	query := `SELECT code, name FROM departments WHERE code = $1`

	row := r.db.QueryRow(context.Background(), query, code)
	department := &models.Department{}
	err := row.Scan(&department.Code, &department.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("department not found")
		}
		return nil, err
	}

	return department, nil
}

func (r *DepartmentRepository) Update(code string, dpt models.Department) error {
	return nil
}

func (r *DepartmentRepository) Delete(code string) error {
	query := `DELETE FROM departments WHERE code = $1`

	result, err := r.db.Exec(context.Background(), query, code)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("department not found")
	}

	return nil
}
