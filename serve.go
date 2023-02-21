package main

import (
  "net/http"
  "fmt"
)

const PORT = 1337
var fstr = fmt.Sprintf

func main() {
  fmt.Printf("Listening on port %d...\n", PORT)
  http.ListenAndServe(fstr(":%d", PORT), http.FileServer(http.Dir(".")))
}
