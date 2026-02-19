package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type EmployeeService interface {
	CreateEmployee(req *CreateEmployeeRequest) (*Employee, error)
	ListEmployees() ([]*Employee, error)
	ReadEmployee(id uuid.UUID) (*Employee, error)
	UpdateEmployee(id uuid.UUID, req *UpdateEmployeeRequest) (*Employee, error)
	DeleteEmployee(id uuid.UUID) error
}

type EmployeeRepository interface {
	Create(employee *Employee) error
	List() ([]*Employee, error)
	Read(id uuid.UUID) (*Employee, error)
	Update(employee *Employee) error
	Delete(id uuid.UUID) error
}

type Employee struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Sex            string    `json:"sex"`
	Age            int       `json:"age"`
	Salary         int       `json:"salary"`
	DepartmentCode *string   `json:"department_code"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateEmployeeRequest struct {
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Age            int     `json:"age"`
	Salary         int     `json:"salary"`
	DepartmentCode *string `json:"department_code"`
}

type UpdateEmployeeRequest struct {
	Name           *string `json:"name"`
	Sex            *string `json:"sex"`
	Age            *int    `json:"age"`
	Salary         *int    `json:"salary"`
	DepartmentCode *string `json:"department_code"`
}
