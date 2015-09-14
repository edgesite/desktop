// build +windows

package desktop

import (
  "testing"
  "unsafe"
)

func Test(t *testing.T) {
	g := GUIDNew("374DE290-123F-4565-9164-39C4925E467B")
	b := []byte{ 144, 226, 77, 55, 63, 18, 101, 69, 145, 100, 57, 196, 146, 94, 70, 123,  }
	if len(g.data) != len(b) {
			t.Error("bad len")
	}

	for i := range(b) {
		if g.data[i] != b[i] {
			t.Error("bad byte array")
		}
	}
	
	if uint32(unsafe.Sizeof(g)) != 16 {
			t.Error("bad size")
	}

}
