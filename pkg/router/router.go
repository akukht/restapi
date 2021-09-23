package router

import (
	"log"
	"net/http"
	"restapi/pkg/controller"

	"github.com/gorilla/mux"
)

func Router() {
	// Initialization mux router
	r := mux.NewRouter()

	//Events requests
	r.Handle("/api/events/filter", controller.IsAuthorized(controller.FilterEvents)).Methods("GET")
	r.Handle("/api/events/filter/week", controller.IsAuthorized(controller.WeekFilterEvents)).Methods("GET")
	r.Handle("/api/events/filter/between", controller.IsAuthorized(controller.BetweenFilterEvents)).Methods("GET")
	r.Handle("/api/events", controller.IsAuthorized(controller.GetEvents)).Methods("GET")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.GetEvent)).Methods("GET")
	r.Handle("/api/events", controller.IsAuthorized(controller.CreateEvent)).Methods("POST")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.UpdateEvent)).Methods("PUT")
	r.Handle("/api/events/{id}", controller.IsAuthorized(controller.DeleteEvent)).Methods("DELETE")

	//User requests
	r.Handle("/api/users", controller.IsAuthorized(controller.UpdateUser)).Methods("PUT")

	//Authorization requests
	r.HandleFunc("/api/login", controller.Login).Methods("POST")
	r.Handle("/api/logout", controller.IsAuthorized(controller.Logout)).Methods("POST")

	//Port listening
	log.Fatal(http.ListenAndServe(":8000", r))
}
