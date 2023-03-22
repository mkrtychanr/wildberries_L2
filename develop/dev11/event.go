package main

import "time"

type Event struct {
	EventName string    `json:"event_name"`
	UserId    string    `json:"user_id"`
	Date      string    `json:"date"`
	Time      time.Time `json:"-"`
}
