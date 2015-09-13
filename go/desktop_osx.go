// +build darwin

package desktop

import (
	"github.com/nfnt/resize"
	"image"
	"os"
)

// user application data folder
func getAppDataFolder() string {
	return path(NSApplicationSupportDirectory, NSUserDomainMask)
}

// user home "/home/user"
func getHomeFolder() string {
	return os.Getenv("HOME")
}

// user my documents "~/Documents"
func getDocumentsFolder() string {
	return path(NSDocumentDirectory, NSUserDomainMask)
}

// user downloads "~/Downloads"
func getDownloadsFolder() string {
	return path(NSDownloadsDirectory, NSUserDomainMask)
}

// user desktop "~/Desktop"
func getDesktopFolder() string {
	return path(NSDesktopDirectory, NSUserDomainMask)
}

func path(d int, dd int) string {
	f := NSFileManagerNew()
	defer f.Release()

	a := f.URLsForDirectoryInDomains(d, dd)
	defer a.Release()

	if a.Count() != 1 {
		return ""
	}

	var u NSURL = NSURLPointer(a.ObjectAtIndex(0))
	defer u.Release()

	return u.Path()
}

//
// SysTray
//

type DesktopSysTrayOSX struct {
	statusbar  NSStatusBar
	statusitem NSStatusItem
	image      NSImage
}

func desktopSysTrayNew() *DesktopSysTray {
	return &DesktopSysTray{os: &DesktopSysTrayOSX{statusbar: NSStatusBarSystemStatusBar()}}
}

func update(m *DesktopSysTray) {
	var d *DesktopSysTrayOSX = m.os.(*DesktopSysTrayOSX)

	if d.statusitem.Pointer == nil {
		d.statusitem = d.statusbar.StatusItemWithLength(NSVariableStatusItemLength)
	}

	if m.Menu != nil {
		mn := createSubMenu(m.Menu)
		mn.SetAutoenablesItems(false)
		d.statusitem.SetMenu(mn)
	}

	d.statusitem.SetToolTip(m.Title)
	d.statusitem.SetHighlightMode(true)
	d.statusitem.SetImage(d.image)

	//Runtime_Loop()
}

func desktopMain() {
  Runtime_Main()
}

func createSubMenu(mm []Menu) NSMenu {
	var mn NSMenu

	for _, m := range mm {
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
			mn.AddItem(menu)
		case MenuSeparator:
			menu := NSMenuItemSeparatorItem()
			defer menu.Release()
			mn.AddItem(menu)
		case MenuCheckBox:
		}
	}

	return mn
}

func setIcon(m *DesktopSysTray, icon image.Image) {
	var d *DesktopSysTrayOSX = m.os.(*DesktopSysTrayOSX)

	if d.image.Pointer != nil {
		d.image.Release()
		d.image.Pointer = nil
	}

	d.image = convertTrayIcon(icon)
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

func show(m *DesktopSysTray) {
	update(m)
}

func hide(m *DesktopSysTray) {
	close(m)
}

func close(m *DesktopSysTray) {
	var d *DesktopSysTrayOSX = m.os.(*DesktopSysTrayOSX)

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
