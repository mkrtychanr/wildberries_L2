package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (s *Server) EventByName(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	eventName, ok := m["event_name"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return

	}
	s.Mu.RLock()
	event, ok := s.Cache[eventName[0]]
	s.Mu.RUnlock()
	if !ok {
		makeJsonRespond(w, 500, jsonError("no event for this name"))
		return
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	day, ok := m["day"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(day[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour)
	timeTo := timeFrom.AddDate(0, 0, 1)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	week, ok := m["week"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(week[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour * 7)
	timeTo := timeFrom.AddDate(0, 0, 7)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	month, ok := m["month"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(month[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if event.Time.Month() == time.Month(handeledTime) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}
