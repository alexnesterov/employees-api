package main

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed templates/*
var templatesFS embed.FS

func main() {
	port := getEnv("PORT", "3000")
	apiURL := getEnv("API_URL", "http://localhost:8080")

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"deref": func(s *string) string {
			if s == nil {
				return ""
			}
			return *s
		},
	}).ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		log.Fatal("failed to parse templates:", err)
	}

	client := NewAPIClient(apiURL)
	h := NewUIHandler(client, tmpl)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", h.Index)
	mux.HandleFunc("GET /employees/table", h.EmployeesTable)
	mux.HandleFunc("POST /employees", h.CreateEmployee)
	mux.HandleFunc("GET /employees/{id}/edit", h.EditEmployee)
	mux.HandleFunc("PUT /employees/{id}", h.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", h.DeleteEmployee)
	mux.HandleFunc("POST /departments", h.CreateDepartment)
	mux.HandleFunc("DELETE /departments/{id}", h.DeleteDepartment)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting UI server on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down UI server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("UI server stopped")
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
