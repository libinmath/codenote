package main
import (
  "fmt"
  "time"
)
// goroutine not exit
// if not stack, may main exit first
func main() {
  go func() {
    i:=0
    for {
      time.Sleep(time.Second)
      fmt.Printf("%d\n",i)
      i++
    }
    fmt.Println("it works")
  }()
  
  msgs := make(chan string)
  <-msgs

}
