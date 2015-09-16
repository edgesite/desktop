// +build windows

package desktop

import (
	"image"
	"github.com/nfnt/resize"
	"unsafe"
	"encoding/base64"
	"strings"
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

func decodeImageString(s string) image.Image {
	i, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(s)))
	if err != nil {
		panic(err)
	}
	return i
}

func convertMenuIcon(i image.Image) *BitmapImage {
	menubarHeigh := getSystemMenuImageSize()

	c := resize.Resize(menubarHeigh, menubarHeigh, i, resize.Lanczos3)

	return BitmapImageNew(c)
}

func getSystemMenuImageSize() uint {
	return uint(UINTPtr(GetSystemMetrics.Call(Arg(SM_CYMENUCHECK))))
}

func getSystemMenuFont() HFONT {
	nm := NONCLIENTMETRICS{}
	nm.cbSize = UINT(unsafe.Sizeof(nm))

	SystemParametersInfo.Call(Arg(SPI_GETNONCLIENTMETRICS), NULL, Arg(&nm), NULL)

	return HFONTPtr(CreateFontIndirect.Call(Arg(&nm.lfMenuFont)))
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
		w = WNDPROCNew(m.DefWindowProc)
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

func (m *Window) DefWindowProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULTPtr(DefWindowProc.Call(Arg(hWnd), Arg(msg), Arg(wParam), Arg(lParam)))
}

func (m *Window) Close() {
	m.Wnd.Close()
	m.WndClassEx.Close()
}

//
// MenuWin
//

type MenuItemWin struct {
	Menu *Menu
	Image *BitmapImage
}

func (m *MenuItemWin) Close() {
	m.Image.Close()
}

//
// DesktopSysTrayWin
// 

type DesktopSysTrayWin struct {
	MainMenu    HMENU
	MenuItems []*MenuItemWin
	Icon    HICON
	MainWnd *Window

	Checked_png *BitmapImage
	Unchecked_png *BitmapImage
}

func desktopSysTrayNew() *DesktopSysTray {
	d := &DesktopSysTrayWin{}
	m := &DesktopSysTray{os: d}

	d.MainWnd = WindowNew(WNDPROCNew(m.WndProc))

	d.Checked_png = convertMenuIcon(decodeImageString(checked_png))
	d.Unchecked_png = convertMenuIcon(decodeImageString(unchecked_png))

	return m
}

func (m *DesktopSysTray) WndProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	switch msg {
	case WM_SHELLNOTIFY:
		switch lParam {
		case WM_LBUTTONUP:
			return LRESULT(0)
		case WM_LBUTTONDBLCLK:
			return LRESULT(0)
		case WM_RBUTTONUP:
			m.showContextMenu()
			return LRESULT(0)
		}
	case WM_COMMAND:
		i := int(wParam)
		mn := d.MenuItems[i]
		if mn.Menu.Action != nil {
			mn.Menu.Action(mn.Menu)
		}
		return LRESULT(0)
	case WM_MEASUREITEM:
		ms := MEASUREITEMSTRUCTPtr(uintptr(lParam))

		i := int(ms.itemData)
		mn := d.MenuItems[i]
		
		hdc := HDCPtr(GetDC.Call(Arg(d.MainWnd.Wnd)))
		defer ReleaseDC.Call(Arg(d.MainWnd.Wnd), Arg(hdc))
		font := getSystemMenuFont()
		defer font.Close()
		fontold := HFONTPtr(SelectObject.Call(Arg(hdc), Arg(font)))
		size := SIZE{}
		GetTextExtentPoint32.Call(Arg(hdc), Arg(mn.Menu.Name), Arg(len(mn.Menu.Name)), Arg(&size))
		SelectObject.Call(Arg(hdc), Arg(fontold))
		size.cx += LONG(getSystemMenuImageSize() + SPACE_ICONS) * 2
		ms.itemWidth = UINT(size.cx)
		ms.itemHeight = UINT(size.cy)
		return LRESULT(0)
	case WM_DRAWITEM:
		di := (*DRAWITEMSTRUCT)(unsafe.Pointer(lParam))
		
		i := int(di.itemData)
		mn := d.MenuItems[i]
		
		if !mn.Menu.Enabled {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_GRAYTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENU)))))
		} else if ((di.itemState & ODS_SELECTED) == ODS_SELECTED) {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_HIGHLIGHTTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_HIGHLIGHT)))))
		} else {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENUTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENU)))))
		}
		
		x := di.rcItem.left
		y := di.rcItem.top
		
		x += LONG(getSystemMenuImageSize() + SPACE_ICONS) * 2;
		
		font := getSystemMenuFont()
		defer font.Close()
		SelectObject.Call(Arg(di.hDC), Arg(font))
		ExtTextOut.Call(Arg(di.hDC), Arg(x), Arg(y), Arg(ETO_OPAQUE), Arg(&di.rcItem), Arg(mn.Menu.Name), Arg(len(mn.Menu.Name)), NULL)

		x = di.rcItem.left
		
		if mn.Menu.Type == MenuCheckBox {
			if mn.Menu.State  {
                d.Checked_png.Draw(x, y, di.hDC)
			} else {
                d.Unchecked_png.Draw(x, y,di.hDC)
			}
		}

        x += LONG(getSystemMenuImageSize() + SPACE_ICONS)
        if (mn.Image != nil) {
            mn.Image.Draw(x, y, di.hDC)
        }
	case WM_QUIT:
		PostMessage.Call(Arg(d.MainWnd.Wnd), Arg(WM_QUIT), NULL, NULL)
	}

	if msg == WM_TASKBARCREATED {
		m.show()
		return LRESULT(0)
	}

	return d.MainWnd.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (m *DesktopSysTray) setIcon(i image.Image) {
	d := m.os.(*DesktopSysTrayWin)

	bm := BitmapImageNew(i)
	defer bm.Close()

	if d.Icon != 0 {
		d.Icon.Close()
	}
	d.Icon = HICONNew(bm.hbm)
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

	if d.MainMenu != 0 {
		d.MainMenu.Close()
		d.MainMenu = 0
	}
}

func (m *DesktopSysTray) showContextMenu() {
	//d := m.os.(*DesktopSysTrayWin)

	m.updateMenus()
}

func (m *DesktopSysTray) updateMenus() {
	d := m.os.(*DesktopSysTrayWin)

    if d.MainMenu != 0 {
		d.MainMenu.Close()
	}

    d.MainMenu = m.createSubMenu(m.Menu)
	
	var pos POINT
	if !BOOLPtr(GetCursorPos.Call(Arg(&pos))).Bool() {
		panic(GetLastErrorString())
	}
	
	for !BOOLPtr(TrackPopupMenu.Call(Arg(d.MainMenu), TPM_RIGHTBUTTON, Arg(pos.x), Arg(pos.y), NULL, Arg(d.MainWnd.Wnd), NULL)).Bool() {
		var hWnd HWND
		// 0x000005a6 - "Popup menu already active."
		if LastError == 0x000005a6 {
			for {
                // "#32768" - pop up menu window class
				hWnd = HWNDPtr(FindWindowEx.Call(NULL, Arg(hWnd), Arg("#32768"), NULL))
				if hWnd == 0 {
					break
				}
				SendMessage.Call(Arg(hWnd), Arg(WM_KEYDOWN), Arg(VK_ESCAPE), NULL)
			}
            // noting is working...
            // just return.
			return
		} else {
			panic(GetLastErrorString())
		}
	}
}

func (m *DesktopSysTray) createSubMenu(mm []Menu) HMENU {
	d := m.os.(*DesktopSysTrayWin)

	hmenu := HMENUPtr(CreatePopupMenu.Call())
	if hmenu == 0 {
		panic(GetLastErrorString())
	}

	for i := range mm {
		mn := &mm[i]

		switch mn.Type {
		case MenuItem, MenuCheckBox:
			menuwin := &MenuItemWin{}
			menuwin.Menu = mn

			if mn.Icon != nil {
				menuwin.Image = convertMenuIcon(mn.Icon)
			}

			id := len(d.MenuItems)
			d.MenuItems = append(d.MenuItems, menuwin)

			if mn.Menu != nil {
				sub := m.createSubMenu(mn.Menu)
		        // seems like you dont have to free this menu, since it already attached
		        // to main HMENU handler
				if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MF_POPUP | MFT_OWNERDRAW), Arg(sub), NULL)).Bool() {
					panic(GetLastErrorString())
				}
				mi := MENUITEMINFO{}
				mi.cbSize = UINT(unsafe.Sizeof(mi))
				if !BOOLPtr(GetMenuItemInfo.Call(Arg(hmenu), Arg(sub), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
				mi.dwItemData = ULONG_PTR(id)
				mi.fMask |= MIIM_DATA
				if !BOOLPtr(SetMenuItemInfo.Call(Arg(hmenu), Arg(sub), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
			} else {
				if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MFT_OWNERDRAW), Arg(id), NULL)).Bool() {
					panic(GetLastErrorString())
				}
				mi := MENUITEMINFO{}
				mi.cbSize = UINT(unsafe.Sizeof(mi))
				if !BOOLPtr(GetMenuItemInfo.Call(Arg(hmenu), Arg(id), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
				mi.dwItemData = ULONG_PTR(id)
				mi.fMask |= MIIM_DATA
				if !BOOLPtr(SetMenuItemInfo.Call(Arg(hmenu), Arg(id), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
			}
		case MenuSeparator:
			if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MF_SEPARATOR), NULL, NULL)).Bool() {
				panic(GetLastErrorString())
			}
		}
	}

	return hmenu
}
