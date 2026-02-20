package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/domain/departments/model"
)

type DepartmentHandler struct {
	svc model.DepartmentService
}

func NewDepartmentHandler(svc model.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		svc: svc,
	}
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req model.CreateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	createdDepartment, err := h.svc.CreateDepartment(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to create department",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Department created successfully",
		Data:    createdDepartment,
	})
}

func (h *DepartmentHandler) ListDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, err := h.svc.ListDepartments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to get departments",
			Details: err.Error(),
		})
		return
	}

	if len(list) == 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(SuccessResponse{
			Message: "No departments found",
			Data:    list,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Departments retrieved successfully",
		Data:    list,
	})
}

func (h *DepartmentHandler) ReadDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	department, err := h.svc.ReadDepartment(id)
	if errors.Is(err, model.ErrDepartmentNotFound) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Department not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to get department",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Department retrieved successfully",
		Data:    department,
	})
}

func (h *DepartmentHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	err := h.svc.DeleteDepartment(id)
	if errors.Is(err, model.ErrDepartmentNotFound) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Department not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Failed to delete department",
			Details: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Department deleted successfully",
	})
}
