package repository

import (
	"context"
	"errors"

	"github.com/alexnesterov/employees-api/internal/model"
	"github.com/jackc/pgx/v5"
)

type employeePgRepo struct {
	db *pgx.Conn
}

func NewEmployeePgRepo(db *pgx.Conn) model.EmployeeRepo {
	return &employeePgRepo{
		db: db,
	}
}

func (r *employeePgRepo) Create(e *model.Employee) error {
	query := `INSERT INTO employees (name, sex, age, salary, department_code) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	params := []any{e.Name, e.Sex, e.Age, e.Salary, e.DepartmentCode}

	row := r.db.QueryRow(context.Background(), query, params...)
	err := row.Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *employeePgRepo) List() ([]*model.Employee, error) {
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

func (r *employeePgRepo) Read(id string) (*model.Employee, error) {
	query := `SELECT id, name, sex, age, salary, department_code FROM employees WHERE id = $1`

	row := r.db.QueryRow(context.Background(), query, id)
	employee := &model.Employee{}
	err := row.Scan(&employee.ID, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary, &employee.DepartmentCode)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return employee, nil
}

func (r *employeePgRepo) Update(id string, e model.Employee) error {
	query := `UPDATE employees SET name = $1, sex = $2, age = $3, salary = $4, department_code = $5 WHERE id = $6`

	_, err := r.db.Exec(context.Background(), query, e.Name, e.Sex, e.Age, e.Salary, e.DepartmentCode, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *employeePgRepo) Delete(id string) error {
	query := `DELETE FROM employees WHERE id = $1`

	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *employeePgRepo) UpdateDepartment(e *model.Employee) error {
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
