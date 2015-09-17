// +build darwin

package desktop

import (
	"github.com/nfnt/resize"
	"image"
)

func desktopMain() {
	Runtime_Main()
}

func convertTrayIcon(i image.Image) NSImage {
	var f NSFont = NSFontMenuBarFontOfSize(0)
	defer f.Release()

	fd := f.FontDescriptor()
	defer fd.Release()

	n := NSNumberPointer(fd.ObjectForKey(NSFontSizeAttribute))
	defer n.Release()

	menubarHeigh := uint(n.IntValue())

	c := resize.Resize(menubarHeigh, menubarHeigh, i, resize.Lanczos3)

	return NSImageImage(c)
}

func convertMenuIcon(i image.Image) NSImage {
	var f NSFont = NSFontMenuFontOfSize(0)
	defer f.Release()

	fd := f.FontDescriptor()
	defer fd.Release()

	n := NSNumberPointer(fd.ObjectForKey(NSFontSizeAttribute))
	defer n.Release()

	menubarHeigh := uint(n.IntValue())

	c := resize.Resize(menubarHeigh, menubarHeigh, i, resize.Lanczos3)

	return NSImageImage(c)
}

func createSubMenu(mm []Menu) NSMenu {
	var mn NSMenu = NSMenuNew()

	for i := range mm {
		m := &mm[i]

		switch m.Type {
		case MenuItem:
			var icon NSImage
			var menu NSMenuItem = NSMenuItemNew()
			if m.Icon != nil {
				icon = convertMenuIcon(m.Icon)
				defer icon.Release()
			}
			menu.SetTitle(m.Name)
			menu.SetImage(icon)
			menu.SetEnabled(m.Enabled)

			if m.Action != nil {
				a := DesktopSysTrayActionOSXNew(m)
				menu.SetTarget(a.Pointer)
				menu.SetAction(DesktopSysTrayActionOSXAction)
			}

			if m.Menu != nil {
				sub := createSubMenu(m.Menu)
				defer sub.Release()
				menu.SetSubmenu(sub)
			}
			mn.AddItem(menu)
		case MenuSeparator:
			menu := NSMenuItemSeparatorItem()
			defer menu.Release()
			mn.AddItem(menu)
		case MenuCheckBox:
			var icon NSImage
			var menu NSMenuItem = NSMenuItemNew()
			if m.Icon != nil {
				icon = convertMenuIcon(m.Icon)
				defer icon.Release()
			}
			menu.SetTitle(m.Name)
			menu.SetImage(icon)
			menu.SetEnabled(m.Enabled)
			menu.SetState((map[bool]int{true: NSOnState, false: NSOffState})[m.State])

			if m.Action != nil {
				a := DesktopSysTrayActionOSXNew(m)
				menu.SetTarget(a.Pointer)
				menu.SetAction(DesktopSysTrayActionOSXAction)
			}

			mn.AddItem(menu)
		}
	}

	mn.SetAutoenablesItems(false)
	return mn
}

type DesktopSysTrayOSX struct {
	statusbar  NSStatusBar
	statusitem NSStatusItem
	image      NSImage
}

func desktopSysTrayNew() *DesktopSysTray {
	return &DesktopSysTray{os: &DesktopSysTrayOSX{statusbar: NSStatusBarSystemStatusBar()}}
}

func (m *DesktopSysTray) update() {
	d := m.os.(*DesktopSysTrayOSX)

	if d.statusitem.Pointer == nil {
		d.statusitem = d.statusbar.StatusItemWithLength(NSVariableStatusItemLength)
	}

	if m.Menu != nil {
		mn := createSubMenu(m.Menu)
		defer mn.Release()
		d.statusitem.SetMenu(mn)
	}

	d.statusitem.SetToolTip(m.Title)
	d.statusitem.SetHighlightMode(true)
	d.statusitem.SetImage(d.image)
}

func (m *DesktopSysTray) setIcon(icon image.Image) {
	d := m.os.(*DesktopSysTrayOSX)

	if d.image.Pointer != nil {
		d.image.Release()
		d.image.Pointer = nil
	}

	d.image = convertTrayIcon(icon)
}

func (m *DesktopSysTray) show() {
	m.update()
}

func (m *DesktopSysTray) hide() {
	m.close()
}

func (m *DesktopSysTray) close() {
	d := m.os.(*DesktopSysTrayOSX)

	if d.statusitem.Pointer != nil {
		d.statusbar.RemoveStatusItem(d.statusitem)
		d.statusitem.Release()
		d.statusitem.Pointer = nil
	}

	if d.statusbar.Pointer != nil {
		d.statusbar.Release()
		d.statusbar.Pointer = nil
	}

	if d.image.Pointer != nil {
		d.image.Release()
		d.image.Pointer = nil
	}
}
