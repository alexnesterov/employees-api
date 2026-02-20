package service

import (
	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
)

type employeeService struct {
	employeeRepo model.EmployeeRepository
}

func NewEmployeeService(repo model.EmployeeRepository) model.EmployeeService {
	return &employeeService{
		employeeRepo: repo,
	}
}

func (s *employeeService) CreateEmployee(req *model.CreateEmployeeRequest) (*model.Employee, error) {
	employee := &model.Employee{
		ID:             uuid.New(),
		Name:           req.Name,
		Sex:            req.Sex,
		Age:            req.Age,
		Salary:         req.Salary,
		DepartmentCode: req.DepartmentCode,
	}

	if err := s.employeeRepo.Create(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *employeeService) ListEmployees() ([]*model.Employee, error) {
	return s.employeeRepo.List()
}

func (s *employeeService) ReadEmployee(id uuid.UUID) (*model.Employee, error) {
	return s.employeeRepo.Read(id)
}

func (s *employeeService) UpdateEmployee(id uuid.UUID, req *model.UpdateEmployeeRequest) (*model.Employee, error) {
	employee, err := s.employeeRepo.Read(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		employee.Name = *req.Name
	}

	if req.Sex != nil {
		employee.Sex = *req.Sex
	}

	if req.Age != nil {
		employee.Age = *req.Age
	}

	if req.Salary != nil {
		employee.Salary = *req.Salary
	}

	if req.DepartmentCode != nil {
		employee.DepartmentCode = req.DepartmentCode
	}

	if err := s.employeeRepo.Update(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *employeeService) DeleteEmployee(id uuid.UUID) error {
	return s.employeeRepo.Delete(id)
}
