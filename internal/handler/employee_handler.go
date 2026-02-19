package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/domain/employees/model"
	"github.com/google/uuid"
)

type EmployeeHandler struct {
	employeeSvc model.EmployeeService
}

func NewEmployeeHandler(svc model.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeSvc: svc,
	}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	createdEmployee, err := h.employeeSvc.CreateEmployee(&req)
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

	var req model.UpdateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	updatedEmployee, err := h.employeeSvc.UpdateEmployee(id, &req)
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

	if err := h.employeeSvc.DeleteEmployee(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to delete employee",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Employee deleted successfully",
	})
}
