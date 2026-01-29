package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexnesterov/employees-api/internal/model"
)

type employeeMemRepository struct {
	counter int
	data    map[string]model.Employee
	sync.Mutex
}

func NewEmployeeMemRepository() model.EmployeeRepository {
	return &employeeMemRepository{
		data:    make(map[string]model.Employee),
		counter: 1,
	}
}

func (r *employeeMemRepository) Create(e *model.Employee) error {
	r.Lock()
	defer r.Unlock()

	e.ID = fmt.Sprint(r.counter)
	r.data[string(e.ID)] = *e
	r.counter++

	return nil
}

func (r *employeeMemRepository) List() ([]*model.Employee, error) {
	r.Lock()
	defer r.Unlock()

	listEmployee := make([]*model.Employee, 0)

	for _, val := range r.data {
		listEmployee = append(listEmployee, &val)
	}

	return listEmployee, nil
}

func (r *employeeMemRepository) Read(id string) (*model.Employee, error) {
	r.Lock()
	defer r.Unlock()

	employee, ok := r.data[id]
	if !ok {
		return &employee, errors.New("employee not found")
	}

	return &employee, nil
}

func (r *employeeMemRepository) Update(id string, e model.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = e

	return nil
}

func (r *employeeMemRepository) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	delete(r.data, id)

	return nil
}

func (r *employeeMemRepository) UpdateDepartment(e *model.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[e.ID] = *e

	return nil
}
