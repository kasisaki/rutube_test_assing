package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	Db.Create(&employee)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var user User
	if err := Db.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var subscription Subscription
	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	subscription.UserID = user.ID
	Db.Create(&subscription)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscription)
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var user User
	if err := Db.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	subscriptionID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	Db.Delete(&Subscription{}, subscriptionID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Unsubscribed"})
}
