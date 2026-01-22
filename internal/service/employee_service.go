package service

import "github.com/alexnesterov/employees-api/internal/model"

type EmployeeService struct {
	repo model.EmployeeRepo
}

func NewEmployeeService(repo model.EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (s *EmployeeService) CreateEmployee(e *model.Employee) error {
	return s.repo.Create(e)
}

func (s *EmployeeService) ListEmployees() ([]*model.Employee, error) {
	return s.repo.List()
}

func (s *EmployeeService) ReadEmployee(id string) (*model.Employee, error) {
	return s.repo.Read(id)
}

func (s *EmployeeService) UpdateEmployee(id string, e model.Employee) error {
	return s.repo.Update(id, e)
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	return s.repo.Delete(id)
}
