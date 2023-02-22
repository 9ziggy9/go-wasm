package main

import (
  // "bufio"
  "context"
  "net/http"
  "fmt"
  // "os"
)

var fstr = fmt.Sprintf

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
  // Instantiate a server struct
  sv := &http.Server{Addr: ":1337"}

  go runServer(sv)

  fmt.Scanln() // Wait for input

  fmt.Println("Killing server.")
  if err := sv.Shutdown(context.Background()); err != nil {
    fmt.Printf("Error stopping server! \n%v\n", err)
  }

  fmt.Println("Server stopped.")
}
