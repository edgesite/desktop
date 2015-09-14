// +build windows

package desktop

import (
	"fmt"
	"os"
)

func getAppDataFolder() string {
	return path(CSIDL_LOCAL_APPDATA)
}

func getHomeFolder() string {
	return os.Getenv("USERPROFILE")
}

func getDocumentsFolder() string {
	return path(CSIDL_PERSONAL)
}

func getDownloadsFolder() string {
	if IsWindowsXP() {
		// xp has no default downloads folder. so be it ~/Documents :)
		return getDocumentsFolder()
	} else {
		// vista+ has a download folder
		//
		// http://stackoverflow.com/questions/7672774/how-do-i-determine-the-windows-download-folder-path
		//
		var guid GUID = GUIDNew("374DE290-123F-4565-9164-39C4925E467B")
		return knowpath(guid)
	}
}

func getDesktopFolder() string {
	return path(CSIDL_DESKTOPDIRECTORY)
}

func knowpath(guid GUID) string {
	var pszPath uintptr
	hResult, _, _ := SHGetKnownFolderPath.Call(Ptr(&guid), Ptr(SHGFP_TYPE_CURRENT), Ptr(0), Ptr(&pszPath))
	switch hResult {
	case S_FILE_NOT_FOUND:
		panic("File not Found")
	case S_OK:
		path := WString2String(pszPath)
		CoTaskMemFree.Call(pszPath)
		return path
	default:
		panic(fmt.Sprint("Error: %x", hResult))
	}
}

func path(nFolder int) string {
	pszPath := [MAX_PATH]uint16{}
	hResult, _, _ := SHGetFolderPath.Call(Ptr(0), Ptr(nFolder), Ptr(0), Ptr(SHGFP_TYPE_CURRENT), Ptr(&pszPath[0]))
	if S_OK == hResult {
		return WString2String(Ptr(&pszPath[0]))
	} else {
		panic(fmt.Sprint("Error: %x", hResult))
	}
}
