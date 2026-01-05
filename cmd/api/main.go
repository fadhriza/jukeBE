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
	// 1. Load Configuration
	cfg := config.LoadConfig()
	log.Printf("Starting application in %s mode", cfg.AppEnv)

	// 2. Connect to Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// 3. Setup Layers (Dependency Injection)
	userRepo := repository.NewUserRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	userService := service.NewUserService(userRepo)
	employeeService := service.NewEmployeeService(employeeRepo)
	authService := service.NewAuthService(adminRepo)

	userHandler := handler.NewUserHandler(userService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	healthHandler := handler.NewHealthHandler(db)
	authMiddleware := middleware.CreateBasicAuthMiddleware(&authService)


	// 4. Setup Router
	mux := http.NewServeMux()

	// Public Routes
	mux.HandleFunc("GET /api/v1/health", healthHandler.Check)
	
	// Static Files - /public/photoprofile/...
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Protected Routes (Basic Auth)
	// User
	mux.Handle("GET /api/v1/user", authMiddleware(http.HandlerFunc(userHandler.GetUser)))

	// Employee Endpoints
	mux.Handle("GET /api/employees", authMiddleware(http.HandlerFunc(employeeHandler.GetAll)))
	mux.Handle("GET /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.GetOne)))
	mux.Handle("POST /api/employees", authMiddleware(http.HandlerFunc(employeeHandler.Create)))
	mux.Handle("PUT /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.Update)))
	mux.Handle("DELETE /api/employees/{id}", authMiddleware(http.HandlerFunc(employeeHandler.Delete)))

	// 5. Start Server
	log.Printf("Server listening on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
