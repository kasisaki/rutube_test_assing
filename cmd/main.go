package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/kasisaki/rutube_test_assing/internal"
	"net/http"
	"time"
)

func main() {
	router := chi.NewRouter()

	router.Post("/login", internal.Login)

	router.Group(func(r chi.Router) {
		r.Use(internal.Authenticate)
		r.Post("/employees", internal.AddEmployee)
		r.Post("/subscribe", internal.Subscribe)
		r.Delete("/unsubscribe/{id}", internal.Unsubscribe)
	})

	go func() {
		for {
			internal.SendBirthdayNotifications()
			time.Sleep(24 * time.Hour)
		}
	}()

	http.ListenAndServe(":8080", router)
}
