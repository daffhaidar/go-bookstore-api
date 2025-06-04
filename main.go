package main

import (
	"log"
	"net/http"

	_ "bookstore/docs"
	"bookstore/handlers"
	"bookstore/middleware"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Bookstore API
// @version 1.0
// @description This is a sample bookstore server.
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// Apply logging middleware
	r.Use(middleware.LoggingMiddleware)

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Routes
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")

	log.Printf("Server starting on port 8080...")
	log.Printf("Swagger UI available at http://localhost:8080/swagger/")
	log.Fatal(http.ListenAndServe(":8080", r))
} 