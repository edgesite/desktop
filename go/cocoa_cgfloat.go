// +build darwin

package desktop

import (
  "unsafe"
  "math"
)

func CGFloatPointer(p unsafe.Pointer) float64 {
  i := uint64(uintptr(p))
  return math.Float64frombits(i)
}
