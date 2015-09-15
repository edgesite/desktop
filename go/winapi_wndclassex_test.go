// build +windows

package desktop

import (
	"testing"
	"unsafe"
)

func TestWNDCLASSEX(t *testing.T) {
	m := &WNDCLASSEX{}
	m.cbSize = UINT(unsafe.Sizeof(*m))

	if m.cbSize != 80 {
		t.Error("wrong WNDCLASSEX size", m.cbSize)
	}
}
