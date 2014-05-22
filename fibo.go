package main

import "fmt"

func fibo(c chan int) {
  a, b := 1, 1
  for {
    a, b = b, a + b
    c <- b
  }
}

func main() {
  c := make(chan int)
  go fibo(c)
  for i := 0; i < 20; i++ {
    fmt.Println(<-c)
  }
}
