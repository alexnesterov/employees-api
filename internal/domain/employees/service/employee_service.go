package service

import (
	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
)

type employeeRepository interface {
	Create(e *model.Employee) (*model.Employee, error)
	List() ([]*model.Employee, error)
	Read(id uuid.UUID) (*model.Employee, error)
	Update(id uuid.UUID, e *model.Employee) (*model.Employee, error)
	Delete(id uuid.UUID) error
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

func (s *EmployeeService) CreateEmployee(e *model.Employee) (*model.Employee, error) {
	return s.employeeRepo.Create(e)
}

func (s *EmployeeService) ListEmployees() ([]*model.Employee, error) {
	return s.employeeRepo.List()
}

func (s *EmployeeService) ReadEmployee(id uuid.UUID) (*model.Employee, error) {
	return s.employeeRepo.Read(id)
}

func (s *EmployeeService) UpdateEmployee(id uuid.UUID, e *model.Employee) (*model.Employee, error) {
	return s.employeeRepo.Update(id, e)
}

func (s *EmployeeService) DeleteEmployee(id uuid.UUID) error {
	return s.employeeRepo.Delete(id)
}

func (s *EmployeeService) UpdateEmployeeDepartment(e *model.Employee) error {
	return s.employeeRepo.UpdateDepartment(e)
}
