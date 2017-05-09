package main

import (
  "fmt"
  "time"
)

func from1(c1 chan string) {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
}
func from2(c2 chan string) {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second *2 )
    }
  }
func reciever(c1,c2 chan string) {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      }
    }
  }
func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go from1(c1)

  go from2(c2)

  go reciever(c1,c2)

  var input string
  fmt.Scanln(&input)
}
