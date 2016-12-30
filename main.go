package main

import (
  "fmt"
  "encoding/json"
)

func main() {
  u1 := &User{Login: "toto"}
  u2 := &User{Login: "titi"}

  m1 := NewMessage(u1, "Je suis toto")
  m2 := NewMessage(u2, "Je suis titi")

  r1 := &Room{Name: "General", Messages: []Message{m1, m2}}

  err := r1.save()
  if err != nil {
    panic(err)
  }

  r2, err := loadRoom("General")
  if err != nil {
    panic(err)
  }

  fmt.Println(json.Marshal(r2.Messages))
}
