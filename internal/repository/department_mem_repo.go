package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexnesterov/employees-api/internal/model"
)

type departmentMemRepo struct {
	counter int
	data    map[string]*model.Department
	sync.Mutex
}

func NewDepartmentMemRepo() model.DepartmentRepo {
	return &departmentMemRepo{
		data:    make(map[string]*model.Department),
		counter: 1,
	}
}

func (r *departmentMemRepo) Create(d *model.Department) error {
	r.Lock()
	defer r.Unlock()

	if d.Code == "" {
		d.Code = fmt.Sprintf("dept-%d", r.counter)
		r.counter++
	}

	r.data[d.Code] = d
	return nil
}

func (r *departmentMemRepo) List() ([]*model.Department, error) {
	r.Lock()
	defer r.Unlock()

	departments := make([]*model.Department, 0, len(r.data))
	for _, dept := range r.data {
		departments = append(departments, dept)
	}

	return departments, nil
}

func (r *departmentMemRepo) Read(code string) (*model.Department, error) {
	r.Lock()
	defer r.Unlock()

	dept, exists := r.data[code]
	if !exists {
		return nil, errors.New("department not found")
	}

	return dept, nil
}

func (r *departmentMemRepo) Update(code string, d model.Department) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.data[code]; !exists {
		return errors.New("department not found")
	}

	d.Code = code
	r.data[code] = &d
	return nil
}

func (r *departmentMemRepo) Delete(code string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.data[code]; !exists {
		return errors.New("department not found")
	}

	delete(r.data, code)
	return nil
}
