package console

import (
  "syscall/js"
)

func Log[T js.Value | string](val T) js.Value {
  return js.Global().Get("console").Call("log", val)
}
