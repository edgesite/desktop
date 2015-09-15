// +build windows

package desktop

import (
	"image"
)

const (
	WM_LBUTTONDOWN   = 513
	WM_NCCREATE      = 129
	WM_NCCALCSIZE    = 131
	WM_CREATE        = 1
	WM_SIZE          = 5
	WM_MOVE          = 3
	WM_USER          = 1024
	WM_LBUTTONUP     = 0x0202
	WM_LBUTTONDBLCLK = 515
	WM_RBUTTONUP     = 517
	WM_CLOSE         = 0x0010
	WM_NULL          = 0x0000
	SW_SHOW          = 5
	WM_COMMAND       = 0x0111
	WM_SHELLNOTIFY   = WM_USER + 1
	WM_MEASUREITEM   = 44
	WM_DRAWITEM      = 43
	WM_CANCELMODE    = 0x001F
	VK_ESCAPE        = 0x1B
	WM_KEYDOWN       = 0x0100
	WM_KEYUP         = 0x0101

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
	Menu HMENU
	Icon HICON
}

func desktopSysTrayNew() *DesktopSysTray {
	m := &DesktopSysTray{os: &DesktopSysTrayWin{}}

	if MainWnd == nil {
		MainWnd = MessageLoopNew()
	}

	return m
}

func setIcon(m *DesktopSysTray, i image.Image) {
	//var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

}

func show(m *DesktopSysTray) {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	n.SetTooltip(m.Title)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_ADD), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
}

func hide(m *DesktopSysTray) {
}

func update(m *DesktopSysTray) {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	d.Menu = HMENUPtr(CreatePopupMenu.Call())
}

func close(m *DesktopSysTray) {
}
