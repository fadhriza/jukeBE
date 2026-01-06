package main

import (
	"log"
	"net/http"

	"jukeBE/config"
	"jukeBE/internal/handler"
	"jukeBE/internal/repository"
	"jukeBE/internal/service"
	"jukeBE/pkg/database"
	"jukeBE/pkg/middleware"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting application in %s mode", cfg.AppEnv)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	employeeRepo := repository.NewEmployeeRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	employeeService := service.NewEmployeeService(employeeRepo)
	authService := service.NewAuthService(adminRepo)

	employeeHandler := handler.NewEmployeeHandler(employeeService)
	healthHandler := handler.NewHealthHandler(db)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := middleware.CreateBasicAuthMiddleware(authService)


	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", healthHandler.Check)
	mux.HandleFunc("POST /api/v1/login", authHandler.Login)
	mux.HandleFunc("POST /api/v1/logout", authHandler.Logout)
	mux.HandleFunc("GET /api/v1/checksession", authHandler.CheckSession)
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))


	mux.Handle("GET /api/employees", authMiddleware(http.HandlerFunc(employeeHandler.GetAll)))
	mux.Handle("GET /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.GetOne)))
	mux.Handle("POST /api/employees", authMiddleware(http.HandlerFunc(employeeHandler.Create)))
	mux.Handle("PUT /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.Update)))
	mux.Handle("DELETE /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.Delete)))

	log.Printf("Server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
