package main

import (
	"log"
	"net/http"
	"simple-backend/config"
	"simple-backend/controllers"
	"simple-backend/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Подключение к базе данных
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Закрываем подключение к БД после завершения
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Создаем новый роутер
	r := mux.NewRouter()

	// Создаем подмаршрутизатор для защищенных маршрутов
	protected := r.PathPrefix("/items").Subrouter()

	// Настройка маршрутов без проверки токена
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/items", controllers.GetItems).Methods("GET")

	// Настройка защищенных маршрутов
	protected.HandleFunc("/create", controllers.CreateItem).Methods("POST")
	protected.HandleFunc("/{id:[0-9]+}", controllers.UpdateItem).Methods("PUT", "PATCH")
	protected.HandleFunc("/{id:[0-9]+}", controllers.DeleteItem).Methods("DELETE")

	// Применяем JWT middleware только к защищенным маршрутам
	protected.Use(middleware.JWTAuth)

	// Добавляем CORS middleware
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Запуск сервера с CORS
	log.Println("Server is starting on port 8080...")
	err = http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)) // Используем r (mux router), а не http.DefaultServeMux
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
