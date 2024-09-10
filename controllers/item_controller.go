package controllers

import (
	"encoding/json"
	"net/http"
	"simple-backend/models"
	"simple-backend/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := repositories.GetAllItems()
	if err != nil {
		http.Error(w, "Error retrieving items", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item

	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := repositories.CreateItem(&newItem); err != nil {
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Извлекаем id из URL с помощью mux.Vars
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Для PUT или PATCH обрабатываем по-разному
	if r.Method == "PUT" {
		// Полное обновление
		var updatedItem models.Item
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := repositories.UpdateItem(uint(id), &updatedItem); err != nil {
			http.Error(w, "Error updating item", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedItem)

	} else if r.Method == "PATCH" {
		// Частичное обновление
		updatedFields := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&updatedFields); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := repositories.PartialUpdateItem(uint(id), updatedFields); err != nil {
			http.Error(w, "Error partially updating item", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedFields)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Извлекаем id из URL с помощью mux.Vars
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Удаляем элемент с указанным ID
	if err := repositories.DeleteItem(uint(id)); err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item deleted successfully"))
}
