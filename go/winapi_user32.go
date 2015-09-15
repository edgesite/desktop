// build +windows

package desktop

import (
	"syscall"
)

const (
	WS_OVERLAPPED           = 0
	WS_OVERLAPPEDWINDOW     = 0x00cf0000
	SPI_GETNONCLIENTMETRICS = 0x0029
	COLOR_MENU              = 4
	COLOR_MENUTEXT          = 7
	COLOR_HIGHLIGHTTEXT     = 14
	COLOR_HIGHLIGHT         = 13
	COLOR_GRAYTEXT          = 17
)

var User32Dll = syscall.MustLoadDLL("User32.dll")
var CreatePopupMenu = User32Dll.MustFindProc("CreatePopupMenu")
var RegisterClassEx = User32Dll.MustFindProc("RegisterClassExW")
var UnregisterClass = User32Dll.MustFindProc("UnregisterClassW")
var CreateWindowEx = User32Dll.MustFindProc("CreateWindowExW")
var GetMessage = User32Dll.MustFindProc("GetMessageW")
var DispatchMessage = User32Dll.MustFindProc("DispatchMessageW")
var DefWindowProc = User32Dll.MustFindProc("DefWindowProcW")
var DestroyWindow = User32Dll.MustFindProc("DestroyWindow")
var CreateIconIndirect = User32Dll.MustFindProc("CreateIconIndirect")
var GetDC = User32Dll.MustFindProc("GetDC")
var ReleaseDC = User32Dll.MustFindProc("ReleaseDC")
var DestroyMenu = User32Dll.MustFindProc("DestroyMenu")
var GetWindowText = User32Dll.MustFindProc("GetWindowTextW")
var GetClassName = User32Dll.MustFindProc("GetClassNameW")
var SetWindowText = User32Dll.MustFindProc("SetWindowTextW")
var TranslateMessage = User32Dll.MustFindProc("TranslateMessage")
