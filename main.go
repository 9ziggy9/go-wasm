package main

import (
  "fmt"
  "syscall/js"
  . "gogol/interfaces"
  // . "gogol/global"
  // "gogol/console"
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

type GridStruct struct {
  cols int
  rows int
  scale float64
  containerId string
  gridId string
  cellClass string
} 

func main() {
  gridStruct := GridStruct{
    cols: 32,
    rows: 18,
    scale: 50,
    containerId: "game",
    gridId: "grid",
    cellClass: "cell",
  }

  gridParams := js.ValueOf(map[string] interface{} {
    "cols": gridStruct.cols,
    "rows": gridStruct.rows,
    "scale": gridStruct.scale,
    "containerId": gridStruct.containerId,
    "gridId": gridStruct.gridId,
    "cellClass": gridStruct.cellClass,
  })

  gridBinding := js.Global().Get("Grid").Invoke(gridParams)
  gridBinding.Get("create").Invoke()
  // Initialize UI
  // _initGrid(60,30)

  // Build the world
  var world World
  world.Build()

  select {} // Run indefinitely
}
