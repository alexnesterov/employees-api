package repository

import (
	"errors"
	"sync"

	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
)

type MemEmployeeRepository struct {
	data map[uuid.UUID]model.Employee
	sync.Mutex
}

func NewMemEmployeeRepository() *MemEmployeeRepository {
	return &MemEmployeeRepository{
		data: make(map[uuid.UUID]model.Employee),
	}
}

func (r *MemEmployeeRepository) Create(e *model.Employee) error {
	r.Lock()
	defer r.Unlock()

	e.ID = uuid.New()
	r.data[e.ID] = *e

	return nil
}

func (r *MemEmployeeRepository) List() ([]*model.Employee, error) {
	r.Lock()
	defer r.Unlock()

	listEmployee := make([]*model.Employee, 0)

	for _, val := range r.data {
		listEmployee = append(listEmployee, &val)
	}

	return listEmployee, nil
}

func (r *MemEmployeeRepository) Read(id uuid.UUID) (*model.Employee, error) {
	r.Lock()
	defer r.Unlock()

	employee, ok := r.data[id]
	if !ok {
		return &employee, errors.New("employee not found")
	}

	return &employee, nil
}

func (r *MemEmployeeRepository) Update(id uuid.UUID, e model.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = e

	return nil
}

func (r *MemEmployeeRepository) Delete(id uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	delete(r.data, id)

	return nil
}

func (r *MemEmployeeRepository) UpdateDepartment(e *model.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[e.ID] = *e

	return nil
}
