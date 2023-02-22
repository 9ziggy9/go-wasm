package main

import (
  // "bufio"
  "context"
  "net/http"
  "fmt"
  // "os"
)

var fstr = fmt.Sprintf

func buildServer(port uint16) *http.Server {
  sv := new(http.Server)
  fmt.Println("Building server structure...")
  fileServer := http.FileServer(http.Dir("."))
  sv = &http.Server{Addr: fstr(":%d", port), Handler: http.DefaultServeMux}
  http.Handle("/", fileServer)
  return sv
}

func runServer(server *http.Server) {
  fmt.Println("Starting server...")
  fmt.Println("Press enter to kill server...")
  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      fmt.Printf("Error starting server: ")
      panic(err)
  }
  defer func() {
    fmt.Println("Adios...")
  }()
}

func main() {
  var input string // Store command line input

  sv := buildServer(1337)
  go runServer(sv)

  fmt.Scanln(&input) // Wait for input
  fmt.Println("You said:", input)

  fmt.Println("Killing server.")
  if err := sv.Shutdown(context.Background()); err != nil {
    fmt.Printf("Error stopping server! \n%v\n", err)
  }

  fmt.Println("Server stopped.")
}
