package service

import "github.com/alexnesterov/employees-api/internal/domain/departments/model"

type departmentService struct {
	repo model.DepartmentRepository
}

func NewDepartmentService(repo model.DepartmentRepository) model.DepartmentService {
	return &departmentService{
		repo: repo,
	}
}

func (s *departmentService) CreateDepartment(req *model.CreateDepartmentRequest) (*model.Department, error) {
	department := &model.Department{
		Code: req.Code,
		Name: req.Name,
	}

	if err := s.repo.Create(department); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *departmentService) ListDepartments() ([]*model.Department, error) {
	return s.repo.List()
}

func (s *departmentService) ReadDepartment(code string) (*model.Department, error) {
	return s.repo.Read(code)
}

func (s *departmentService) DeleteDepartment(code string) error {
	return s.repo.Delete(code)
}
