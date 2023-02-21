package main

import (
  "syscall/js"
)

func main() {
  console := js.Global().Get("console")
  console.Call("log", "Hello, Go!")
}
