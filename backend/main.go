package main

import (
	"log"
	"net/http"
)

const SECRET = "very-secret-stuff"

var addr string = ":8080"

var productStorage *ProductStorage
var userStorage []User

func main() {
	db := initDB()
	defer db.Close()
	log.Printf("database connection established, dsn: %s", DB_DSN)

	productStorage = NewProductStorage(db)

	mux := http.NewServeMux()
	manager := NewManager()

	mux.HandleFunc("GET /api/health-check", healthCheck)

	mux.Handle("GET /api/products", manager.With(http.HandlerFunc(getProducts)))
	mux.HandleFunc("POST /api/products", createProduct)
	mux.HandleFunc("PUT /api/products/{id}", updateProduct)
	mux.HandleFunc("DELETE /api/products/{id}", deleteProduct)
	mux.HandleFunc("POST /api/users", createUser)

	server := http.Server{
		Addr:    addr,
		Handler: logger(corsMiddleware(mux)),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println("something went wrong", err)
	}
}
