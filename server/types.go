package main

import "time"

type Message struct {
	Username string
	Content  string
	Date     time.Time
}
