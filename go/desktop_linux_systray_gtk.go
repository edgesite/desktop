// +build linux

package desktop

import (
	"image"
)

type DesktopSysTrayGtk struct {
	m *DesktopSysTray
}

func (os *DesktopSysTrayGtk) show() {
}

func (os *DesktopSysTrayGtk) hide() {
}

func (os *DesktopSysTrayGtk) update() {
}

func (os *DesktopSysTrayGtk) close() {
}

func (os *DesktopSysTrayGtk) setIcon(i image.Image) {
}

