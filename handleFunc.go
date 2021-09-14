package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func betweenFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	var day int

	startNum, _ := strconv.Atoi(strings.Join(urlParams["start"], ""))
	endNum, _ := strconv.Atoi(strings.Join(urlParams["end"], ""))

	var event = map[string]Events{}
	for id, val := range EventsData {
		day, _ = strconv.Atoi(val.Time.Day)

		if day >= startNum && day <= endNum && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}
	json.NewEncoder(w).Encode(event)
}

func weekFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := r.URL.Query()
	weekNum, _ := strconv.Atoi(strings.Join(urlParams["week"], ""))

	weekNum = weekNum * 7
	var day int
	var event = map[string]Events{}
	for id, val := range EventsData {
		day, _ = strconv.Atoi(val.Time.Day)

		if day < weekNum && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}
	json.NewEncoder(w).Encode(event)
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
