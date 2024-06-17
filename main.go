package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func main() {
	router := chi.NewRouter()

	router.Post("/login", login)

	router.Group(func(r chi.Router) {
		r.Use(authenticate)
		r.Post("/employees", addEmployee)
		r.Post("/subscribe", subscribe)
		r.Delete("/unsubscribe/{id}", unsubscribe)
	})

	go func() {
		for {
			sendBirthdayNotifications()
			time.Sleep(24 * time.Hour)
		}
	}()

	http.ListenAndServe(":8080", router)
}
