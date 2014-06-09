package main

import (
  "github.com/thoj/go-ircevent"
  "fmt"
)

func main() {
  con := irc.IRC("my-bot-test", "my-bot-test")
  err := con.Connect("irc.freenode.net:6667")

  if err != nil {
    fmt.Println("Failed connecting")
    return
  }

  con.AddCallback("001", func (e *irc.Event) {
    con.Join("#minasdev")
  })

  con.Loop()
}
