package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Port struct {
	Port string `json:"port"`
}

type Server struct {
	Mu    sync.RWMutex
	Cache map[string]Event
	port  Port
}

func newServer() (*Server, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	port := Port{}
	err = json.Unmarshal(data, &port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cache: make(map[string]Event),
		port:  port,
	}, nil
}

func (s *Server) setEvent(event Event) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Cache[event.EventName] = event
}

func (s *Server) deleteEvent(eventName string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Cache, eventName)
}

func (s *Server) setupGetHandlers() {
	http.HandleFunc("/event_by_name", MiddlewareLogger(s.EventByName))
	http.HandleFunc("/events_for_day", MiddlewareLogger(s.EventsForDay))
	http.HandleFunc("/events_for_week", MiddlewareLogger(s.EventsForWeek))
	http.HandleFunc("/events_for_month", MiddlewareLogger(s.EventsForMonth))
}

func (s *Server) setupPostHandlers() {
	http.HandleFunc("/create_event", MiddlewareLogger(s.CreateEvent))
	http.HandleFunc("/update_event", MiddlewareLogger(s.UpdateEvent))
	http.HandleFunc("/delete_event", MiddlewareLogger(s.DeleteEvent))
}

func (s *Server) SetupHandlers() {
	s.setupGetHandlers()
	s.setupPostHandlers()
}

func (s *Port) getAddress() string {
	return fmt.Sprintf(":%s", s.Port)
}

func (s *Server) Up() {
	address := s.port.getAddress()
	fmt.Println("Server listen on", address)
	log.Println(http.ListenAndServe(address, nil))
}
