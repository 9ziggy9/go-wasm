package main

import (
  "fmt"
  "syscall/js"
  . "gogol/interfaces"
  "gogol/ui"
)

// NOTE:
// You need to create a new scope for each js.FuncOf closure, so that the values of x and y
// are captured at the time the closure is created. One way to do this is to define a new
// function that takes x and y as arguments and returns the closure. 
// TODO: Generalize this mechanism

func handleClick(x, y int) js.Func {
  return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    fmt.Println(fmt.Sprintf("(%d,%d)", x, y))
    return nil
  })
}

func main() {
  // Build the world
  var world World
  world.Build()

  // Bind to JS UI
  gridBinding := ui.GenerateGridBinding(world, 50, "game", "grid", "cell")
  gridBinding.Get("create").Invoke()

  select {} // Run indefinitely
}
