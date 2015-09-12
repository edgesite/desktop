// +build darwin

package desktop

import (
	"os"
  "image"
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
  statusitem NSStatusItem
  image NSImage
}

func desktopSysTrayNew() *DesktopSysTray {
	return &DesktopSysTray{os: &DesktopSysTrayOSX{}}
}

func update(m *DesktopSysTray) {
	var d *DesktopSysTrayOSX = m.os.(*DesktopSysTrayOSX)
  
  statusbar := NSStatusBarSystemStatusBar()
  defer statusbar.Release()
  
  if d.statusitem.Pointer == nil {
    d.statusitem = statusbar.StatusItemWithLength(NSVariableStatusItemLength)
  }
  
  d.statusitem.SetToolTip(m.Title)
  d.statusitem.SetHighlightMode(true)
  d.statusitem.SetImage(d.image)
}

func setIcon(m *DesktopSysTray, icon image.Image) {
	var d *DesktopSysTrayOSX = m.os.(*DesktopSysTrayOSX)
  
  d.image = NSImageImage(icon)
}

func show(m *DesktopSysTray) {
  update(m)
}

func hide(m *DesktopSysTray) {
}

func close(m *DesktopSysTray) {
}
