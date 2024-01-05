package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	ListenAddr string
	Messages   []Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Messages:   []Message{},
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/sendmessage", s.sendMessageHandler)
	http.HandleFunc("/getmessage", s.getMessageHandler)
	fmt.Println("running on: " + s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, nil)
}
