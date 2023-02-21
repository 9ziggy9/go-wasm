package main

import (
  "syscall/js"
)

func _consoleLog(str string) js.Value {
  return js.Global().Get("console").Call("log", str)
}

func main() {
  _consoleLog("Hello, Go! Lel")
}
