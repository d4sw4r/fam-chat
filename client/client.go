package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Uri      string
	Username string
}

var contenttype string = "application/json"

func NewClient(uri, username string) *Client {
	return &Client{Uri: uri, Username: username}
}

func (c *Client) sendmessage(msg Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	bodyReader := bytes.NewReader(body)
	http.Post(c.Uri+"/sendmessage", contenttype, bodyReader)
	return err

}

func (c *Client) getmessage() ([]Message, error) {
	req, err := http.Get(c.Uri + "/getmessage")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		fmt.Println(err)
	}
	return messages, err
}
