package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	permissionError     int = 2
	valid               int = 1
	invalidData         int = 0
	internalServerError int = -1
)

func getDataFromRequest(r *http.Request) (Event, error) {
	event := Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return event, err
	}
	err = json.Unmarshal(data, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (s *Server) validatePost(w http.ResponseWriter, event *Event, actionType string) int {
	s.Mu.RLock()
	data, ok := s.Cache[event.EventName]
	s.Mu.RUnlock()
	result := invalidData

	date := isValidDate(event)
	switch actionType {
	case "create":
		if !ok && date && event.EventName != "" {
			result = valid
		}
	case "update":
		if ok && date {
			result = valid
		}
	case "delete":
		if ok {
			if data.UserId == event.UserId {
				result = valid
			} else {
				result = permissionError
			}
		}
	default:
		result = internalServerError
	}
	return result
}

func (s *Server) validateAndRespond(w http.ResponseWriter, code int) bool {
	if code == valid {
		return true
	}
	switch code {
	case internalServerError:
		makeJsonRespond(w, 503, jsonError("internal server error"))
	case invalidData:
		makeJsonRespond(w, 400, jsonError("invalid data"))
	case permissionError:
		makeJsonRespond(w, 500, jsonError("permisson error"))
	}
	return false
}

func (s *Server) postRequestCheck(w http.ResponseWriter, r *http.Request, request string) (Event, error) {
	event := Event{}
	if r.Method != http.MethodPost {
		errorString := "method not allowed"
		makeJsonRespond(w, 500, jsonError(errorString))
		return event, fmt.Errorf(errorString)
	}
	event, err := getDataFromRequest(r)
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return event, err
	}
	validate := s.validatePost(w, &event, request)
	if !s.validateAndRespond(w, validate) {
		return event, fmt.Errorf("something being wrong")
	}
	return event, nil
}

func (s *Server) createAndUpdate(w http.ResponseWriter, r *http.Request, request string) {
	event, err := s.postRequestCheck(w, r, request)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.setEvent(event)
	makeJsonRespond(w, 200, jsonResult("ok"))
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "create")
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "update")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := s.postRequestCheck(w, r, "delete")
	if err != nil {
		return
	}
	s.deleteEvent(event.EventName)
	makeJsonRespond(w, 200, jsonResult("ok"))
}
