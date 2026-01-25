package main

import (
	"context"
	"log"

	"github.com/alexnesterov/employees-api/internal/config"
	"github.com/alexnesterov/employees-api/internal/handler"
	"github.com/alexnesterov/employees-api/internal/repository"
	"github.com/alexnesterov/employees-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config.Load()

	connConfig, err := pgx.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to parse database config: %v", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	router := gin.Default()

	employeeRepo := repository.NewEmployeePgRepo(conn)
	employeeService := service.NewEmployeeService(employeeRepo)
	handlerEmployee := handler.NewEmployeeHandler(employeeService)

	router.POST("/employees", handlerEmployee.CreateEmployee)
	router.GET("/employees", handlerEmployee.ListEmployee)
	router.GET("/employees/:id", handlerEmployee.GetEmployee)
	router.PUT("/employees/:id", handlerEmployee.UpdateEmployee)
	router.DELETE("/employees/:id", handlerEmployee.DeleteEmployee)
	router.PUT("/employees/:id/department", handlerEmployee.UpdateEmployeeDepartment)

	departmentRepo := repository.NewDepartmentPgRepo(conn)
	departmentService := service.NewDepartmentService(departmentRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	router.POST("/departments", departmentHandler.CreateDepartment)
	router.GET("/departments", departmentHandler.ListDepartments)
	router.GET("/departments/:id", departmentHandler.ReadDepartment)
	router.DELETE("/departments/:id", departmentHandler.DeleteDepartment)

	router.Run(":" + cfg.Port)
}
