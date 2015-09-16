// +build windows

package desktop

func desktopMain() {
	msg := &MSG{}

	for BOOLPtr(GetMessage.Call(Arg(msg), NULL, NULL, NULL)).Bool() {
		TranslateMessage.Call(Arg(msg))
		DispatchMessage.Call(Arg(msg))
	}
}
