// build +windows

package desktop

import (
	"testing"
)

func TestNOTIFYICONDATA(t *testing.T) {
	n := NOTIFYICONDATANew()

	if n.cbSize != 976 {
		t.Error("wrong NOTIFYICONDATA size", n.cbSize)
	}
}
