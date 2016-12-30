package main

import (
  "time"
)

type Message struct {
  Author *User
  Content string
  Date time.Time
}

func NewMessage(author *User, content string) Message {
  return Message{Author: author, Content: content, Date: time.Now()}
}
