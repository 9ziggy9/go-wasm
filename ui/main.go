package ui

import (
  "syscall/js"
  "gogol/interfaces"
)

func GenerateGridBinding(
  w interfaces.World, scale int, cId string, gId string, cC string,
) js.Value {
  gridParams := js.ValueOf(map[string] interface{} {
    "cols": len(w.Cells[0]),
    "rows": len(w.Cells),
    "scale": scale,
    "containerId": cId,
    "gridId": gId,
    "cellClass": cC,
  })
  return js.Global().Get("Grid").Invoke(gridParams)
}
