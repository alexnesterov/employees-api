package repository

import (
	"context"
	"errors"

	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PgEmployeeRepository struct {
	db *pgx.Conn
}

func NewPgEmployeeRepository(db *pgx.Conn) *PgEmployeeRepository {
	return &PgEmployeeRepository{
		db: db,
	}
}

func (r *PgEmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	query := `INSERT INTO employees (name, sex, age, salary, department_code) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, sex, age, salary, department_code, created_at, updated_at`
	params := []any{e.Name, e.Sex, e.Age, e.Salary, e.DepartmentCode}

	err := r.db.QueryRow(context.Background(), query, params...).Scan(
		&e.ID,
		&e.Name,
		&e.Sex,
		&e.Age,
		&e.Salary,
		&e.DepartmentCode,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *PgEmployeeRepository) List() ([]*model.Employee, error) {
	query := `SELECT id, name, sex, age, salary, department_code FROM employees`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*model.Employee, error) {
		employee := &model.Employee{}
		err := row.Scan(&employee.ID, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary, &employee.DepartmentCode)
		if err != nil {
			return nil, err
		}
		return employee, nil
	})

	return employees, err
}

func (r *PgEmployeeRepository) Read(id uuid.UUID) (*model.Employee, error) {
	query := `SELECT id, name, sex, age, salary, department_code FROM employees WHERE id = $1`

	row := r.db.QueryRow(context.Background(), query, id)
	employee := &model.Employee{}
	err := row.Scan(&employee.ID, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary, &employee.DepartmentCode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, model.ErrEmployeeNotFound
		}
		return nil, err
	}

	return employee, nil
}

func (r *PgEmployeeRepository) Update(id uuid.UUID, e model.Employee) error {
	query := `UPDATE employees SET name = $1, sex = $2, age = $3, salary = $4, department_code = $5 WHERE id = $6`

	_, err := r.db.Exec(context.Background(), query, e.Name, e.Sex, e.Age, e.Salary, e.DepartmentCode, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgEmployeeRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM employees WHERE id = $1`

	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgEmployeeRepository) UpdateDepartment(e *model.Employee) error {
	query := `UPDATE employees SET department_code = $1 WHERE id = $2`

	result, err := r.db.Exec(context.Background(), query, e.DepartmentCode, e.ID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("employee not found")
	}

	return nil
}
