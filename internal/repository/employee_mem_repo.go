package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexnesterov/employees-api/internal/models"
)

type employeeMemRepo struct {
	counter int
	data    map[string]models.Employee
	sync.Mutex
}

func NewEmployeeMemRepo() models.EmployeeRepo {
	return &employeeMemRepo{
		data:    make(map[string]models.Employee),
		counter: 1,
	}
}

func (r *employeeMemRepo) Create(e *models.Employee) error {
	r.Lock()
	defer r.Unlock()

	e.ID = fmt.Sprint(r.counter)
	r.data[string(e.ID)] = *e
	r.counter++

	return nil
}

func (r *employeeMemRepo) List() ([]*models.Employee, error) {
	r.Lock()
	defer r.Unlock()

	listEmployee := make([]*models.Employee, 0)

	for _, val := range r.data {
		listEmployee = append(listEmployee, &val)
	}

	return listEmployee, nil
}

func (r *employeeMemRepo) Read(id string) (*models.Employee, error) {
	r.Lock()
	defer r.Unlock()

	employee, ok := r.data[id]
	if !ok {
		return &employee, errors.New("employee not found")
	}

	return &employee, nil
}

func (r *employeeMemRepo) Update(id string, e models.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = e

	return nil
}

func (r *employeeMemRepo) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	delete(r.data, id)

	return nil
}
