package model

import "errors"

var (
	ErrDepartmentNotFound = errors.New("department not found")
)

type DepartmentService interface {
	CreateDepartment(req *CreateDepartmentRequest) (*Department, error)
	ListDepartments() ([]*Department, error)
	ReadDepartment(code string) (*Department, error)
	DeleteDepartment(code string) error
}

type DepartmentRepository interface {
	Create(department *Department) error
	List() ([]*Department, error)
	Read(code string) (*Department, error)
	Delete(code string) error
}

type Department struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CreateDepartmentRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
