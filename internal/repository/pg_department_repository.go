package repository

import (
	"context"

	"github.com/alexnesterov/employees-api/internal/domain/departments/model"
	"github.com/jackc/pgx/v5"
)

type departmentPgRepo struct {
	db *pgx.Conn
}

func NewDepartmentPgRepo(db *pgx.Conn) model.DepartmentRepository {
	return &departmentPgRepo{
		db: db,
	}
}

func (r *departmentPgRepo) Create(department *model.Department) error {
	query := `
		INSERT INTO departments (code, name)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(context.Background(), query,
		department.Code,
		department.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *departmentPgRepo) List() ([]*model.Department, error) {
	query := `SELECT code, name FROM departments`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	departments, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*model.Department, error) {
		department := &model.Department{}
		err := row.Scan(&department.Code, &department.Name)
		if err != nil {
			return nil, err
		}
		return department, nil
	})

	return departments, err
}

func (r *departmentPgRepo) Read(code string) (*model.Department, error) {
	query := `SELECT code, name FROM departments WHERE code = $1`

	row := r.db.QueryRow(context.Background(), query, code)
	department := &model.Department{}
	err := row.Scan(&department.Code, &department.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrDepartmentNotFound
		}
		return nil, err
	}

	return department, nil
}

func (r *departmentPgRepo) Delete(code string) error {
	query := `DELETE FROM departments WHERE code = $1`

	result, err := r.db.Exec(context.Background(), query, code)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return model.ErrDepartmentNotFound
	}

	return nil
}
