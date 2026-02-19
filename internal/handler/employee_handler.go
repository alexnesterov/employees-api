package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
)

type employeeService interface {
	CreateEmployee(e *model.Employee) (*model.Employee, error)
	ListEmployees() ([]*model.Employee, error)
	ReadEmployee(id uuid.UUID) (*model.Employee, error)
	UpdateEmployee(id uuid.UUID, e *model.Employee) (*model.Employee, error)
	DeleteEmployee(id uuid.UUID) error
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

func (h *EmployeeHandler) ListEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, err := h.employeeSvc.ListEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to get employees",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Employees retrieved successfully",
		Data:    list,
	})
}

func (h *EmployeeHandler) ReadEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idValue := r.PathValue("id")

	id, err := uuid.Parse(idValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid id",
			Details: err.Error(),
		})
		return
	}

	employee, err := h.employeeSvc.ReadEmployee(id)
	if err != nil {
		if errors.Is(err, model.ErrEmployeeNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error:   "Employee not found",
				Details: err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to get employee",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Employee retrieved successfully",
		Data:    employee,
	})
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idValue := r.PathValue("id")
	id, err := uuid.Parse(idValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid id",
			Details: err.Error(),
		})
		return
	}

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	updatedEmployee, err := h.employeeSvc.UpdateEmployee(id, &employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to update employee",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Employee updated successfully",
		Data:    updatedEmployee,
	})
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee deleted"))
}

func (h *EmployeeHandler) UpdateEmployeeDepartment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("employee department updated"))
}
