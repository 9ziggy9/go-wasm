package main

import (
  // "bufio"
  "net/http"
  "fmt"
  // "os"
)

var fstr = fmt.Sprintf

func serveHTTP(port uint16) {
  fmt.Printf("Listening on port %d...\n", port)
  err := http.ListenAndServe(fstr(":%d", port), http.FileServer(http.Dir(".")))
  if err != nil {
    panic("Panicking!")
  }
  defer func() {
    fmt.Println("Goodbye!")
  }()
}

func main() {
  serveHTTP(1337)
}
