package main

import (
	"log"
	"net/http"
	"simple-backend/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/items", handlers.GetItems)
	http.HandleFunc("/items/create", handlers.CreateItem)
	http.HandleFunc("/items/update", handlers.UpdateItem)
	http.HandleFunc("/items/delete", handlers.DeleteItem)

	log.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
