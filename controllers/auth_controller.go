package controllers

import (
	"encoding/json"
	"net/http"
	"simple-backend/repositories"
	"simple-backend/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	username := request["username"]
	password := request["password"]

	user, err := repositories.GetUserByUsername(username)
	if err != nil || user == nil || !user.CheckPassword(password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
