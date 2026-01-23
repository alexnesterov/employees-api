package handler

import (
	"fmt"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/model"
	"github.com/alexnesterov/employees-api/internal/service"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type EmployeeHandler struct {
	svc *service.EmployeeService
}

func NewEmployeeHandler(svc *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		svc: svc,
	}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee model.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.svc.CreateEmployee(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": employee.ID,
	})
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee model.Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.svc.UpdateEmployee(id, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employee, err := h.svc.ReadEmployee(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) ListEmployee(c *gin.Context) {
	list, err := h.svc.ListEmployees()
	if err != nil {
		fmt.Printf("failed to get employee %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(list)

	c.JSON(http.StatusOK, list)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	h.svc.DeleteEmployee(id)

	c.String(http.StatusOK, "employee deleted")
}

func (h *EmployeeHandler) UpdateEmployeeDepartment(c *gin.Context) {
	var employee model.Employee

	employee.ID = c.Param("id")

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.svc.UpdateEmployeeDepartment(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id":              employee.ID,
		"department_code": employee.DepartmentCode,
	})
}
