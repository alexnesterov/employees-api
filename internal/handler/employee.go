package handler

import (
	"fmt"
	"net/http"

	"github.com/alexnesterov/employees-api/internal/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type EmployeeHandler struct {
	repo models.EmployeeRepo
}

func NewEmployeeHandler(repo models.EmployeeRepo) *EmployeeHandler {
	return &EmployeeHandler{
		repo: repo,
	}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.repo.Create(&employee)
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

	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := h.repo.Update(id, employee)
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

	employee, err := h.repo.Read(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) ListEmployee(c *gin.Context) {
	list, err := h.repo.List()
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

	h.repo.Delete(id)

	c.String(http.StatusOK, "employee deleted")
}
