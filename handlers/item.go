package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"simple-backend/models"
)

var items = []models.Item{
	{ID: 1, Name: "Item 1", Description: "Description 1"},
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem models.Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
