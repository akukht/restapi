package router

import (
	"context"
	"log"
	"net/http"
	"restapi/pkg/controller"
	"time"

	"github.com/gorilla/mux"
)

func Router(ctx context.Context) (err error) {

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

	//--------
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return

	//---------
	//Port listening
	//log.Fatal(http.ListenAndServe(":8000", r))

}
