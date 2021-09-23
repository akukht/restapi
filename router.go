package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func router() {
	r := mux.NewRouter()
	r.Handle("/api/events/filter", isAuthorized(filterEvents)).Methods("GET")
	r.Handle("/api/events/filter/week", isAuthorized(weekFilterEvents)).Methods("GET")
	r.Handle("/api/events/filter/between", isAuthorized(betweenFilterEvents)).Methods("GET")
	r.Handle("/api/events", isAuthorized(GetEvents)).Methods("GET")
	r.Handle("/api/events/{id}", isAuthorized(GetEvent)).Methods("GET")
	r.Handle("/api/events", isAuthorized(createEvent)).Methods("POST")
	r.Handle("/api/events/{id}", isAuthorized(updateEvent)).Methods("PUT")
	r.Handle("/api/events/{id}", isAuthorized(DeleteEvent)).Methods("DELETE")

	r.Handle("/api/users", isAuthorized(UpdateUser)).Methods("PUT")

	r.HandleFunc("/api/login", login).Methods("POST")
	r.Handle("/api/logout", isAuthorized(logout)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

}
