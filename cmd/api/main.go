package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexnesterov/employees-api/internal/config"
	"github.com/alexnesterov/employees-api/internal/domain/employees/service"
	"github.com/alexnesterov/employees-api/internal/handler"
	"github.com/alexnesterov/employees-api/internal/repository"

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

	mux := http.NewServeMux()

	employeeRepository := repository.NewPgEmployeeRepository(conn)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	mux.HandleFunc("POST /employees", employeeHandler.CreateEmployee)
	mux.HandleFunc("GET /employees", employeeHandler.ListEmployee)
	// router.GET("/employees/:id", handlerEmployee.GetEmployee)
	// router.PUT("/employees/:id", handlerEmployee.UpdateEmployee)
	// router.DELETE("/employees/:id", handlerEmployee.DeleteEmployee)
	// router.PUT("/employees/:id/department", handlerEmployee.UpdateEmployeeDepartment)

	// departmentRepo := repository.NewDepartmentPgRepo(conn)
	// departmentService := service.NewDepartmentService(departmentRepo)
	// departmentHandler := handler.NewDepartmentHandler(departmentService)

	// router.POST("/departments", departmentHandler.CreateDepartment)
	// router.GET("/departments", departmentHandler.ListDepartments)
	// router.GET("/departments/:id", departmentHandler.ReadDepartment)
	// router.DELETE("/departments/:id", departmentHandler.DeleteDepartment)

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
