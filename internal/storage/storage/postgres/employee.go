package postgres

import (
	"context"
	"errors"

	"github.com/alexnesterov/employees-api/internal/models"
	"github.com/jackc/pgx/v5"
)

type EmployeeRepository struct {
	db *pgx.Conn
}

func NewEmployeeRepository(db *pgx.Conn) models.EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) Create(e *models.Employee) error {
	query := `INSERT INTO employees (name, sex, age, salary) VALUES ($1, $2, $3, $4) RETURNING id`
	params := []any{e.Name, e.Sex, e.Age, e.Salary}

	row := r.db.QueryRow(context.Background(), query, params...)
	err := row.Scan(&e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeRepository) List() ([]*models.Employee, error) {
	query := `SELECT id, name, sex, age, salary FROM employees`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*models.Employee, error) {
		employee := &models.Employee{}
		err := row.Scan(&employee.ID, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary)
		if err != nil {
			return nil, err
		}
		return employee, nil
	})

	return employees, err
}

func (r *EmployeeRepository) Read(id string) (*models.Employee, error) {
	query := `SELECT id, name, sex, age, salary FROM employees WHERE id = $1`

	row := r.db.QueryRow(context.Background(), query, id)
	employee := &models.Employee{}
	err := row.Scan(&employee.ID, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepository) Update(id string, e models.Employee) error {
	query := `UPDATE employees SET name = $1, sex = $2, age = $3, salary = $4 WHERE id = $5`

	_, err := r.db.Exec(context.Background(), query, e.Name, e.Sex, e.Age, e.Salary, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeRepository) Delete(id string) error {
	query := `DELETE FROM employees WHERE id = $1`

	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
