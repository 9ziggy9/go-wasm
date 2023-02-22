package interfaces

// Cell implementation
type CellState int
const (
  DEAD CellState = iota
  ALIVE
)

type Cell struct {
  X uint8
  Y uint8
  State CellState
}
var Kill = func (c *Cell) { c.State = DEAD }
var Birth = func (c *Cell) { c.State = ALIVE }
