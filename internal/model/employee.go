package model

type Employee struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Age            int     `json:"age"`
	Salary         int     `json:"salary"`
	DepartmentCode *string `json:"department_code"`
}

type EmployeeRepository interface {
	Create(e *Employee) error
	List() ([]*Employee, error)
	Read(id string) (*Employee, error)
	Update(id string, e Employee) error
	Delete(id string) error
	UpdateDepartment(e *Employee) error
}

type EmployeeService interface {
	CreateEmployee(e *Employee) error
	ListEmployees() ([]*Employee, error)
	ReadEmployee(id string) (*Employee, error)
	UpdateEmployee(id string, e Employee) error
	DeleteEmployee(id string) error
	UpdateEmployeeDepartment(e *Employee) error
}
