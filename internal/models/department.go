package models

type Department struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type DepartmentRepository interface {
	Create(d *Department) error
	List() ([]*Department, error)
	Read(code string) (*Department, error)
	Update(code string, d Department) error
	Delete(code string) error
}
