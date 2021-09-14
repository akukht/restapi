package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func router() {
	r := mux.NewRouter()
	r.HandleFunc("/api/events/filter", filterEvents).Methods("GET")
	r.HandleFunc("/api/events/filter/week", weekFilterEvents).Methods("GET")
	r.HandleFunc("/api/events/filter/between", betweenFilterEvents).Methods("GET")
	r.HandleFunc("/api/events", getEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", getEvent).Methods("GET")
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events/{id}", updateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
