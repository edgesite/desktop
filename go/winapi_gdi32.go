// build +windows

package desktop

import (
	"syscall"
)

const (
	BI_RGB = 0

	DIB_RGB_COLORS = 0
)

var Gdi32Dll = syscall.MustLoadDLL("Gdi32.dll")
var DeleteObject = Gdi32Dll.MustFindProc("DeleteObject")
var CreateCompatibleDC = Gdi32Dll.MustFindProc("CreateCompatibleDC")
var DeleteDC = Gdi32Dll.MustFindProc("DeleteDC")
var CreateDIBSection = Gdi32Dll.MustFindProc("CreateDIBSection")
