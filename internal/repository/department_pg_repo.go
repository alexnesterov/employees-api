package repository

import (
	"context"
	"errors"

	"github.com/alexnesterov/employees-api/internal/model"
	"github.com/jackc/pgx/v5"
)

type departmentPgRepo struct {
	db *pgx.Conn
}

func NewDepartmentPgRepo(db *pgx.Conn) model.DepartmentRepo {
	return &departmentPgRepo{
		db: db,
	}
}

func (r *departmentPgRepo) Create(dpt *model.Department) error {
	query := `INSERT INTO departments (code, name) VALUES ($1, $2) RETURNING code`
	params := []any{dpt.Code, dpt.Name}

	row := r.db.QueryRow(context.Background(), query, params...)
	err := row.Scan(&dpt.Code)
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
			return nil, errors.New("department not found")
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
		return errors.New("department not found")
	}

	return nil
}
