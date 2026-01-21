package models

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type EmployeeRepo interface {
	Create(e *Employee) error
	List() ([]*Employee, error)
	Read(id string) (*Employee, error)
	Update(id string, e Employee) error
	Delete(id string) error
}
