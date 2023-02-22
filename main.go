package main

import (
  "fmt"
  "syscall/js"
  . "gogol/interfaces"
)

// GLOBAL VARS
const (
  ROWS = 35
  COLS = 50
)

func consoleLog[T js.Value | string](val T) js.Value {
  return js.Global().Get("console").Call("log", val)
}

// Taking the convention that _ means run ONCE
func _initGrid(rows uint8, cols uint8) {
  doc := js.Global().Get("document")
  grid := doc.Call("getElementById", "grid")
  consoleLog(grid)
  for y := 0; y < ROWS; y = y+1 {
    for x := 0; x < COLS; x = x+1 {
      newCell := doc.Call("createElement", "div")
      newCell.Call("setAttribute", "id", fmt.Sprintf("(%d,%d)", x, y))
      grid.Call("appendChild", newCell)
    }
  }
}

func init() {
  _initGrid(60,30)
}

func main() {
  c := Cell{X: 0, Y: 0, State: DEAD}
  fmt.Printf("x: %d, y: %d, state: %d", c.X, c.Y, c.State)
}
