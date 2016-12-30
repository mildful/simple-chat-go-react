package main

import (
  "fmt"
  "encoding/json"
  "net/http"
)

func roomHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Path[len("/room/"):]
  room, err := loadRoom(name)
  if err != nil {
    panic(err)
  }

  content, err := json.Marshal(room.Messages)
  if err != nil {
    panic(err)
  }
  
  fmt.Fprintf(w, "<h1>%s</h1><div>%s<?div>", name, content)
}

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

  http.HandleFunc("/room/", roomHandler)
  http.ListenAndServe(":8089", nil)
}
