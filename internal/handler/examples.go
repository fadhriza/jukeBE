package handler

// examples.go
//
// This package handles HTTP requests. It parses input, calls the Service layer,
// and formats the response (typically JSON).
//
// Example Usage:
//
// handler := handler.NewUserHandler(userService)
// http.HandleFunc("/users", handler.GetUser)
