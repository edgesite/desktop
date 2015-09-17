// +build windows

package desktop

import "runtime"

func desktopMain() {
	msg := &MSG{}

	for BOOLPtr(GetMessage.Call(Arg(msg), NULL, NULL, NULL)).Bool() {
		TranslateMessage.Call(Arg(msg))
		DispatchMessage.Call(Arg(msg))
	}
	
	runtime.UnlockOSThread()
}
