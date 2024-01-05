package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) getMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := json.NewEncoder(w).Encode(s.Messages)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	}

}

func (s *Server) sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg.Date = time.Now()
		if len(s.Messages) > 9 {
			s.Messages = s.Messages[1:]
		}
		s.Messages = append(s.Messages, msg)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
	}
}
