package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/model"
)

type employeeService interface {
	CreateEmployee(e *model.Employee) (*model.Employee, error)
	ListEmployees() ([]*model.Employee, error)
	ReadEmployee(id string) (*model.Employee, error)
	UpdateEmployee(id string, e model.Employee) error
	DeleteEmployee(id string) error
	UpdateEmployeeDepartment(e *model.Employee) error
}

type EmployeeHandler struct {
	employeeSvc employeeService
}

func NewEmployeeHandler(svc employeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeSvc: svc,
	}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	createdEmployee, err := h.employeeSvc.CreateEmployee(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Internal server error",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Employee created successfully",
		Data:    createdEmployee,
	})
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee updated"))
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee"))
}

func (h *EmployeeHandler) ListEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employees"))
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee deleted"))
}

func (h *EmployeeHandler) UpdateEmployeeDepartment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee department updated"))
}
