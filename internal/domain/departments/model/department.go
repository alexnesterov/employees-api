package model

type Department struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type DepartmentRepo interface {
	Create(d *Department) error
	List() ([]*Department, error)
	Read(code string) (*Department, error)
	Delete(code string) error
}
