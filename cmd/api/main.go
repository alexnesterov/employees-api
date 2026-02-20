package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexnesterov/employees-api/internal/config"
	department "github.com/alexnesterov/employees-api/internal/domain/departments/service"
	employee "github.com/alexnesterov/employees-api/internal/domain/employees/service"
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
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	mux.HandleFunc("POST /employees", employeeHandler.CreateEmployee)
	mux.HandleFunc("GET /employees", employeeHandler.ListEmployee)
	mux.HandleFunc("GET /employees/{id}", employeeHandler.ReadEmployee)
	mux.HandleFunc("PUT /employees/{id}", employeeHandler.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", employeeHandler.DeleteEmployee)

	departmentRepo := repository.NewDepartmentPgRepo(conn)
	departmentService := department.NewDepartmentService(departmentRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	mux.HandleFunc("POST /departments", departmentHandler.CreateDepartment)
	mux.HandleFunc("GET /departments", departmentHandler.ListDepartments)
	mux.HandleFunc("GET /departments/{id}", departmentHandler.ReadDepartment)
	mux.HandleFunc("DELETE /departments/{id}", departmentHandler.DeleteDepartment)

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.Printf("Starting server on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
