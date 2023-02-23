package main

import (
  "fmt"
  "syscall/js"
  . "gogol/interfaces"
  . "gogol/global"
  "gogol/console"
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

// Taking the convention that _ means run ONCE
func _initGrid(rows uint8, cols uint8) {
  doc := js.Global().Get("document")
  grid := doc.Call("getElementById", "grid")
  console.Log(grid)
  for y := 0; y < ROWS; y = y+1 {
    for x := 0; x < COLS; x = x+1 {
      newCell := doc.Call("createElement", "div")
      newCell.Call("setAttribute", "id", fmt.Sprintf("(%d,%d)", x, y))
      newCell.Call("addEventListener", "click", handleClick(x,y))
      grid.Call("appendChild", newCell)
    }
  }
}

func main() {
  console.Log("Hello, from Go main!")
  // Initialize UI
  _initGrid(60,30)

  // Build the world
  var world World
  world.Build()
  world.Log()

  select {} // Run indefinitely
}
