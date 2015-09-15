// build +windows

package desktop

import (
	"syscall"
)

const (
	FORMAT_MESSAGE_FROM_SYSTEM = 0x00001000
)

var Kernel32Dll = syscall.MustLoadDLL("Kernel32.dll")
var GetLastError = Kernel32Dll.MustFindProc("GetLastError")
var FormatMessage = Kernel32Dll.MustFindProc("FormatMessageW")
var GetModuleHandle = Kernel32Dll.MustFindProc("GetModuleHandleW")
