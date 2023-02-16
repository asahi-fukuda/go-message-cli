package model

import "time"

type Message struct {
	Name    string
	Message string
	Time    time.Time
}
