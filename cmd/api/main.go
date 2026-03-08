package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexnesterov/employees-api/internal/config"
	department "github.com/alexnesterov/employees-api/internal/domain/departments/service"
	employee "github.com/alexnesterov/employees-api/internal/domain/employees/service"
	"github.com/alexnesterov/employees-api/internal/handler"
	"github.com/alexnesterov/employees-api/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("failed to ping database:", err)
	}
	log.Println("Connected to database")

	mux := http.NewServeMux()

	employeeRepository := repository.NewPgEmployeeRepository(pool)
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	mux.HandleFunc("POST /employees", employeeHandler.CreateEmployee)
	mux.HandleFunc("GET /employees", employeeHandler.ListEmployee)
	mux.HandleFunc("GET /employees/{id}", employeeHandler.ReadEmployee)
	mux.HandleFunc("PUT /employees/{id}", employeeHandler.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", employeeHandler.DeleteEmployee)

	departmentRepository := repository.NewPgDepartmentRepository(pool)
	departmentService := department.NewDepartmentService(departmentRepository)
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
		IdleTimeout:    60 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on port %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Closing database connection...")
	pool.Close()
	log.Println("Server stopped")
}
