package handler

import (
	"net/http"

	"github.com/alexnesterov/employees-api/internal/models"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	repo models.DepartmentRepo
}

func NewDepartmentHandler(repo models.DepartmentRepo) *DepartmentHandler {
	return &DepartmentHandler{
		repo: repo,
	}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var department models.Department

	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := h.repo.Create(&department); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create department",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Department created successfully",
		"data":    department,
	})
}

func (h *DepartmentHandler) ListDepartments(c *gin.Context) {
	list, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get departments",
			"details": err.Error(),
		})
		return
	}

	if len(list) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No departments found",
			"data":    list,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Departments retrieved successfully",
		"data":    list,
	})
}

func (h *DepartmentHandler) ReadDepartment(c *gin.Context) {
	id := c.Param("id")

	department, err := h.repo.Read(id)
	if err != nil {
		if err.Error() == "department not found" {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get department",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Department retrieved successfully",
		"data":    department,
	})
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.Delete(id); err != nil {
		if err.Error() == "department not found" {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete department",
			"details": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
