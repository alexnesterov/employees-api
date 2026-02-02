package service

import "github.com/alexnesterov/employees-api/internal/model"

type employeeRepository interface {
	Create(e *model.Employee) error
	List() ([]*model.Employee, error)
	Read(id string) (*model.Employee, error)
	Update(id string, e model.Employee) error
	Delete(id string) error
	UpdateDepartment(e *model.Employee) error
}

type EmployeeService struct {
	employeeRepo employeeRepository
}

func NewEmployeeService(repo employeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepo: repo,
	}
}

func (s *EmployeeService) CreateEmployee(e *model.Employee) error {
	return s.employeeRepo.Create(e)
}

func (s *EmployeeService) ListEmployees() ([]*model.Employee, error) {
	return s.employeeRepo.List()
}

func (s *EmployeeService) ReadEmployee(id string) (*model.Employee, error) {
	return s.employeeRepo.Read(id)
}

func (s *EmployeeService) UpdateEmployee(id string, e model.Employee) error {
	return s.employeeRepo.Update(id, e)
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	return s.employeeRepo.Delete(id)
}

func (s *EmployeeService) UpdateEmployeeDepartment(e *model.Employee) error {
	return s.employeeRepo.UpdateDepartment(e)
}
