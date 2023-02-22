package main

import (
  "fmt"
  "syscall/js"
  . "gogol/interfaces"
  . "gogol/global"
  "gogol/console"
)

// Taking the convention that _ means run ONCE
func _initGrid(rows uint8, cols uint8) {
  doc := js.Global().Get("document")
  grid := doc.Call("getElementById", "grid")
  console.Log(grid)
  for y := 0; y < ROWS; y = y+1 {
    for x := 0; x < COLS; x = x+1 {
      newCell := doc.Call("createElement", "div")
      newCell.Call("setAttribute", "id", fmt.Sprintf("(%d,%d)", x, y))
      grid.Call("appendChild", newCell)
    }
  }
}

func init() {
  // Initialize UI
  _initGrid(60,30)

  // Build the world
  var world World
  world.Build()
  world.Log()
}

func main() {
  console.Log("Hello, from Go main!")
}
