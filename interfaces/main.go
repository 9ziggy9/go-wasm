package interfaces

import (
  . "gogol/global"
  "gogol/console"
  "fmt"
)

// Cell implementation
type CellState int
const (
  DEAD CellState = iota
  ALIVE
)

type Cell struct {
  X int
  Y int
  State CellState
}
var Kill = func (c *Cell) { c.State = DEAD }
var Birth = func (c *Cell) { c.State = ALIVE }

type World struct {
  Cells [ROWS][COLS]Cell
}
func (w *World) Build() {
  for y := 0; y < ROWS; y = y+1 {
    for x := 0; x < COLS; x = x+1 {
      w.Cells[y][x] = Cell{X: x, Y: y, State: DEAD}
    }
  }
}
func (w *World) Log() {
  worldStr := "["
  for y := 0; y < ROWS; y = y+1 {
    for x := 0; x < COLS; x = x+1 {
      worldStr += fmt.Sprintf("(%d, %d),", w.Cells[y][x].X, w.Cells[y][x].Y)
    }
  }
  worldStr += "]"
  console.Log(worldStr)
}
