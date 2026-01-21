package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexnesterov/employees-api/internal/models"
)

type DepartmentRepository struct {
	counter int
	data    map[string]*models.Department
	sync.Mutex
}

func NewDepartmentRepository() models.DepartmentRepository {
	return &DepartmentRepository{
		data:    make(map[string]*models.Department),
		counter: 1,
	}
}

func (r *DepartmentRepository) Create(d *models.Department) error {
	r.Lock()
	defer r.Unlock()

	if d.Code == "" {
		d.Code = fmt.Sprintf("dept-%d", r.counter)
		r.counter++
	}

	r.data[d.Code] = d
	return nil
}

func (r *DepartmentRepository) List() ([]*models.Department, error) {
	r.Lock()
	defer r.Unlock()

	departments := make([]*models.Department, 0, len(r.data))
	for _, dept := range r.data {
		departments = append(departments, dept)
	}

	return departments, nil
}

func (r *DepartmentRepository) Read(code string) (*models.Department, error) {
	r.Lock()
	defer r.Unlock()

	dept, exists := r.data[code]
	if !exists {
		return nil, errors.New("department not found")
	}

	return dept, nil
}

func (r *DepartmentRepository) Update(code string, d models.Department) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.data[code]; !exists {
		return errors.New("department not found")
	}

	d.Code = code
	r.data[code] = &d
	return nil
}

func (r *DepartmentRepository) Delete(code string) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.data[code]; !exists {
		return errors.New("department not found")
	}

	delete(r.data, code)
	return nil
}
