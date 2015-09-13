// +build darwin

package desktop

import (
	"unsafe"
)

func CGFloatPointer(p unsafe.Pointer) float64 {
	return Pointer2Float(p)
}
