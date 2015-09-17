// build +windows

package desktop

import (
	"syscall"
)

var Ole32Dll = syscall.MustLoadDLL("Ole32.dll")
var CoTaskMemFree = Ole32Dll.MustFindProc("CoTaskMemFree")
