// +build windows

package desktop

import ()

type MessageLoop struct {
	WndClassEx *WNDCLASSEX
	Wnd        HWND
}

func MessageLoopNew() *MessageLoop {
	m := &MessageLoop{}

	hinstance := HINSTANCEPtr(GetModuleHandle.Call())

	m.WndClassEx = WNDCLASSEXNew(hinstance, WNDPROCNew(m.WndProc), "MessageLoop")

	m.Wnd = HWNDPtr(CreateWindowEx.Call(NULL, Arg(m.WndClassEx.lpszClassName), Arg(m.WndClassEx.lpszClassName), Arg(WS_OVERLAPPEDWINDOW), NULL, NULL, NULL, NULL, NULL, NULL, Arg(hinstance), NULL))
	if m.Wnd == 0 {
		panic(GetLastErrorString())
	}

	return m
}

func (m *MessageLoop) WndProc(hWnd HWND, msg int, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULTPtr(DefWindowProc.Call(Arg(hWnd), Arg(msg), Arg(wParam), Arg(lParam)))
}

func (m *MessageLoop) Close() {
	if !BOOLPtr(DestroyWindow.Call(Arg(m.Wnd))).Bool() {
		panic(GetLastErrorString())
	}
	m.WndClassEx.Close()
}

var MainWnd *MessageLoop

func desktopMain() {
	msg := &MSG{}

	for BOOLPtr(GetMessage.Call(Arg(msg), Arg(MainWnd.Wnd), NULL, NULL)).Bool() {
		DispatchMessage.Call(Arg(msg))
	}
}
