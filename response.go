package main

import (
  "log"
	"net/http"
	"encoding/json"
	"bytes"
)

type Response struct {
  Username string `json:"username"`
	Emoji    string `json:"icon_emoji"`
  Channel  string `json:"channel"`
  Text     string `json:"text"`
}

func (r *Response) Send(u string) {
  // Convert the response object into a json string for posting.
  jsonStr, _ := json.Marshal(r)

  // Build a request that will be sent to BrowserStack.
  client := &http.Client{}
  req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonStr))
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")
  _, err = client.Do(req)
  if err != nil {
  	log.Println("Could not post to Slack.")
  }
}
