package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexnesterov/employees-api/internal/models"
)

type EmployeeRepository struct {
	counter int
	data    map[string]models.Employee
	sync.Mutex
}

func NewEmployeeRepository() models.EmployeeRepository {
	return &EmployeeRepository{
		data:    make(map[string]models.Employee),
		counter: 1,
	}
}

func (r *EmployeeRepository) Create(e *models.Employee) error {
	r.Lock()
	defer r.Unlock()

	e.ID = fmt.Sprint(r.counter)
	r.data[string(e.ID)] = *e
	r.counter++

	return nil
}

func (r *EmployeeRepository) List() ([]*models.Employee, error) {
	r.Lock()
	defer r.Unlock()

	listEmployee := make([]*models.Employee, 0)

	for _, val := range r.data {
		listEmployee = append(listEmployee, &val)
	}

	return listEmployee, nil
}

func (r *EmployeeRepository) Read(id string) (*models.Employee, error) {
	r.Lock()
	defer r.Unlock()

	employee, ok := r.data[id]
	if !ok {
		return &employee, errors.New("employee not found")
	}

	return &employee, nil
}

func (r *EmployeeRepository) Update(id string, e models.Employee) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = e

	return nil
}

func (r *EmployeeRepository) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	delete(r.data, id)

	return nil
}
