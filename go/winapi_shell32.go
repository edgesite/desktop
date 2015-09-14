// build +windows

package desktop

import (
	"syscall"
)

const (
	MAX_PATH = 260

	// Local Settings\Application Data
	CSIDL_LOCAL_APPDATA = 0x001c

	// ~/My Documents
	CSIDL_PERSONAL = 0x005

	// ~/Desktop
	CSIDL_DESKTOPDIRECTORY = 0x10

	SHGFP_TYPE_CURRENT = 0
	SHGFP_TYPE_DEFAULT = 1
	S_OK               = 0
	S_FILE_NOT_FOUND   = 0x80070002

	NIM_ADD    = 0
	NIM_MODIFY = 1
	NIM_DELETE = 2
)

var Shell32Dll = syscall.MustLoadDLL("Shell32.dll")
var SHGetKnownFolderPath = Shell32Dll.MustFindProc("SHGetKnownFolderPath")
var SHGetFolderPath = Shell32Dll.MustFindProc("SHGetFolderPathW")
