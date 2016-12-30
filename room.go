package main

import (
  "io/ioutil"
  "encoding/json"
)

type Room struct {
  Name string
  Messages []Message
}

func (r *Room) save() error {
  filename := r.Name + ".txt"
  content, err := json.Marshal(r.Messages)
  if err != nil {
    return err
  }
  return ioutil.WriteFile(filename, content, 0600)
}

func loadRoom(name string) (*Room, error) {
  filename := name + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  var messages []Message
  err = json.Unmarshal(body, &messages)
  if err != nil {
    return nil, err
  }

  return &Room{Name: name, Messages: messages}, nil
}
