// +build linux

package desktop

import (
	"image"
)

type DesktopSysTrayGtk struct {
	m *DesktopSysTray
}

func DesktopSysTrayGtkNew(m *DesktopSysTray) *DesktopSysTrayGtk {
	os := &DesktopSysTrayGtk{m}
	GtkMessageLoopInc()
	return os
}

func (os *DesktopSysTrayGtk) show() {
}

func (os *DesktopSysTrayGtk) hide() {
}

func (os *DesktopSysTrayGtk) update() {
}

func (os *DesktopSysTrayGtk) close() {
	GtkMessageLoopDec()
}

func (os *DesktopSysTrayGtk) setIcon(i image.Image) {
}

