// +build linux

package desktop

import (
	"image"
)

func desktopSysTrayNew() *DesktopSysTray {
	return &DesktopSysTray{}
}

func (m *DesktopSysTray) show() {
}

func (m *DesktopSysTray) hide() {
}

func (m *DesktopSysTray) update() {
}

func (m *DesktopSysTray) close() {
}

func (m *DesktopSysTray) setIcon(i image.Image) {
}

