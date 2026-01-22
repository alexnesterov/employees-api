package service

import "github.com/alexnesterov/employees-api/internal/model"

type DepartmentService struct {
	repo model.DepartmentRepo
}

func NewDepartmentService(repo model.DepartmentRepo) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (s *DepartmentService) CreateDepartment(dpt *model.Department) error {
	return s.repo.Create(dpt)
}

func (s *DepartmentService) ListDepartments() ([]*model.Department, error) {
	return s.repo.List()
}

func (s *DepartmentService) ReadDepartment(code string) (*model.Department, error) {
	return s.repo.Read(code)
}

func (s *DepartmentService) DeleteDepartment(code string) error {
	return s.repo.Delete(code)
}
