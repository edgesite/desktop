// +build windows

package desktop

import (
)

type Window struct {
	WndClassEx *WNDCLASSEX
	Wnd        HWND
	Name WString
}

func WindowNew(w WNDPROC) *Window {
	m := &Window{}

	if w == 0 {
		w = WNDPROCNew(m.DefWindowProc)
	}

	hinstance := HINSTANCEPtr(GetModuleHandle.Call())

	m.WndClassEx = WNDCLASSEXNew(hinstance, w, "SystrayIcon")
	
	m.Name = WStringNew("SystrayIcon Window")

	m.Wnd = HWNDPtr(CreateWindowEx.Call(NULL, Arg(m.WndClassEx.lpszClassName),
		Arg(m.Name), Arg(WS_OVERLAPPEDWINDOW),
		NULL, NULL, NULL, NULL, NULL, NULL, Arg(hinstance), NULL))
	if m.Wnd == 0 {
		panic(GetLastErrorString())
	}

	return m
}

func (m *Window) DefWindowProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULTPtr(DefWindowProc.Call(Arg(hWnd), Arg(msg), Arg(wParam), Arg(lParam)))
}

func (m *Window) Close() {
	m.Wnd.Close()
	m.WndClassEx.Close()
	m.Name.Close()
}