package main

/*
Етап №1:
створити сервіс (REST API) органайзера-калердаря з функціональністю:
- додавати події, нагадування.
- редагувати їх, змінювати назву, час, опис...
- видаляти події
- переглядати перелік подій на день, тиждень, місяць, рік (з пітримкою фільтрації по ознакам)
*/

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type responseCode struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Events struct {
	Name string `json:"name"`
	Time Date   `json:"time"`
	Desc string `json:"desc"`
}

var EventsData = map[string]Events{}

func init() {
	EventsData["1"] = Events{
		Name: "Google Cloud",
		Time: Date{Day: "12", Month: "03", Year: "2021"},
		Desc: "Description Google event",
	}

	EventsData["2"] = Events{
		Name: "Amazon AWS",
		Time: Date{Day: "12", Month: "03", Year: "2021"},
		Desc: "Description Amazon AWS event",
	}

	EventsData["3"] = Events{
		Name: "Microsoft Azure",
		Time: Date{Day: "11", Month: "03", Year: "2022"},
		Desc: "Description Microsoft Azure event",
	}
	EventsData["4"] = Events{
		Name: "Yahoo event",
		Time: Date{Day: "12", Month: "03", Year: "2021"},
		Desc: "Description Yahoo event event",
	}
}

func main() {
	router()
}

func router() {
	r := mux.NewRouter()
	r.HandleFunc("/api/events/filter", filterEvents).Methods("GET")
	r.HandleFunc("/api/events", getEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", getEvent).Methods("GET")
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events/{id}", updateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func filterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := r.URL.Query()
	var event = map[string]Events{}
	for id, val := range EventsData {
		if val.Time.Day == strings.Join(urlParams["day"], "") && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}

	json.NewEncoder(w).Encode(event)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	if _, ok := EventsData[id["id"]]; ok {
		delete(EventsData, id["id"])
	} else {
		err := responseCode{StatusCode: 404, Message: "Event not found"}
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(EventsData)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	var event Events
	_ = json.NewDecoder(r.Body).Decode(&event)

	if _, ok := EventsData[id["id"]]; ok {
		EventsData[id["id"]] = event
	} else {
		err := responseCode{StatusCode: 404, Message: "Event not found"}
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(EventsData)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createEvent Events
	_ = json.NewDecoder(r.Body).Decode(&createEvent)
	EventsData[strconv.Itoa(rand.Intn(1000000))] = createEvent
	json.NewEncoder(w).Encode(EventsData)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	if _, ok := EventsData[id["id"]]; ok {
		json.NewEncoder(w).Encode(EventsData[id["id"]])
	} else {
		err := responseCode{StatusCode: 404, Message: "Event not found"}
		json.NewEncoder(w).Encode(err)
	}
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EventsData)
}
