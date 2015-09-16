// +build windows

package desktop

import (
	"image"
)

var WM_TASKBARCREATED = UINTPtr(RegisterWindowMessage.Call(String2WString("TaskbarCreated")))

const (
	WM_LBUTTONDOWN   DWORD = 513
	WM_NCCREATE            = 129
	WM_NCCALCSIZE          = 131
	WM_CREATE              = 1
	WM_SIZE                = 5
	WM_MOVE                = 3
	WM_USER                = 1024
	WM_LBUTTONUP           = 0x0202
	WM_LBUTTONDBLCLK       = 515
	WM_RBUTTONUP           = 517
	WM_CLOSE               = 0x0010
	WM_NULL                = 0x0000
	SW_SHOW                = 5
	WM_COMMAND             = 0x0111
	WM_SHELLNOTIFY         = WM_USER + 1
	WM_MEASUREITEM         = 44
	WM_DRAWITEM            = 43
	WM_CANCELMODE          = 0x001F
	VK_ESCAPE              = 0x1B
	WM_KEYDOWN             = 0x0100
	WM_KEYUP               = 0x0101

	MF_ENABLED    = 0
	MF_DISABLED   = 0x00000002
	MF_CHECKED    = 0x00000008
	MF_UNCHECKED  = 0
	MF_GRAYED     = 0x00000001
	MF_STRING     = 0x00000000
	MFT_OWNERDRAW = 256
	MF_SEPARATOR  = 0x00000800
	MF_POPUP      = 0x00000010

	TPM_RECURSE     = 0x0001
	TPM_RIGHTBUTTON = 0x0002

	SM_CYMENUCHECK = 72
	SM_CYMENU      = 15

	SPACE_ICONS = 2
)

type DesktopSysTrayWin struct {
	Menu    HMENU
	Icon    HICON
	MainWnd *Window
}

func desktopSysTrayNew() *DesktopSysTray {
	d := &DesktopSysTrayWin{}
	m := &DesktopSysTray{os: d}

	d.MainWnd = WindowNew(WNDPROCNew(m.WndProc))

	return m
}

func (m *DesktopSysTray) WndProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	switch msg {
	case WM_SHELLNOTIFY:
		switch lParam {
		case WM_LBUTTONUP:
		case WM_LBUTTONDBLCLK:
		case WM_RBUTTONUP:
			m.showContextMenu()
		}
	case WM_COMMAND:
	case WM_MEASUREITEM:
	case WM_DRAWITEM:
	case WM_QUIT:
		PostMessage.Call(Arg(d.MainWnd.Wnd), Arg(WM_QUIT), NULL, NULL)
	}

	if msg == WM_TASKBARCREATED {
		m.show()
	}

	return d.MainWnd.WndProc(hWnd, msg, wParam, lParam)
}

func (m *DesktopSysTray) setIcon(i image.Image) {
	d := m.os.(*DesktopSysTrayWin)

	bm := HBITMAPNew(i)
	defer bm.Close()

	if d.Icon != 0 {
		d.Icon.Close()
	}
	d.Icon = HICONNew(bm)
}

func (m *DesktopSysTray) show() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	n.SetTooltip(m.Title)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_ADD), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
}

func (m *DesktopSysTray) hide() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_DELETE), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
}

func (m *DesktopSysTray) update() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	n.SetTooltip(m.Title)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_MODIFY), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
}

func (m *DesktopSysTray) close() {
	d := m.os.(*DesktopSysTrayWin)

	if d.Icon != 0 {
		d.Icon.Close()
		d.Icon = 0
	}

	if d.Menu != 0 {
		d.Menu.Close()
		d.Menu = 0
	}
}

func (m *DesktopSysTray) showContextMenu() {
	//d := m.os.(*DesktopSysTrayWin)

	m.updateMenus()
}

func (m *DesktopSysTray) updateMenus() {
	menu := HMENUPtr(CreatePopupMenu.Call())
	if menu == 0 {
		panic(GetLastErrorString())
	}

}

//
// Window
//

type Window struct {
	WndClassEx *WNDCLASSEX
	Wnd        HWND
}

func WindowNew(w WNDPROC) *Window {
	m := &Window{}

	if w == 0 {
		w = WNDPROCNew(m.WndProc)
	}

	hinstance := HINSTANCEPtr(GetModuleHandle.Call())

	m.WndClassEx = WNDCLASSEXNew(hinstance, w, "MessageLoop")

	m.Wnd = HWNDPtr(CreateWindowEx.Call(NULL, Arg(m.WndClassEx.lpszClassName),
		Arg(m.WndClassEx.lpszClassName), Arg(WS_OVERLAPPEDWINDOW),
		NULL, NULL, NULL, NULL, NULL, NULL, Arg(hinstance), NULL))
	if m.Wnd == 0 {
		panic(GetLastErrorString())
	}

	return m
}

func (m *Window) WndProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULTPtr(DefWindowProc.Call(Arg(hWnd), Arg(msg), Arg(wParam), Arg(lParam)))
}

func (m *Window) Close() {
	m.Wnd.Close()
	m.WndClassEx.Close()
}
