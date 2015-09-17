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
	WM_QUIT                 = 0x0012
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
var RegisterWindowMessage = User32Dll.MustFindProc("RegisterWindowMessageW")
var PostMessage = User32Dll.MustFindProc("PostMessageW")
var AppendMenu = User32Dll.MustFindProc("AppendMenuW")
var GetMenuItemInfo = User32Dll.MustFindProc("GetMenuItemInfoW")
var SetMenuItemInfo = User32Dll.MustFindProc("SetMenuItemInfoW")
var GetSystemMetrics = User32Dll.MustFindProc("GetSystemMetrics")
var SetForegroundWindow = User32Dll.MustFindProc("SetForegroundWindow")
var GetCursorPos = User32Dll.MustFindProc("GetCursorPos")
var TrackPopupMenu = User32Dll.MustFindProc("TrackPopupMenu")
var FindWindowEx = User32Dll.MustFindProc("FindWindowExW")
var SendMessage = User32Dll.MustFindProc("SendMessageW")
var SystemParametersInfo = User32Dll.MustFindProc("SystemParametersInfoW")
var GetSysColor = User32Dll.MustFindProc("GetSysColor")
