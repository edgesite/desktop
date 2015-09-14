// +build windows

package desktop

import (
	"image"
)

func desktopSysTrayNew() *DesktopSysTray {
	return &DesktopSysTray{}
}

func desktopMain(){
}

func setIcon(m *DesktopSysTray, i image.Image) string {
  return ""
}

func show(m *DesktopSysTray) string {
  return ""
}

func hide(m *DesktopSysTray) string {
  return ""
}

func update(m *DesktopSysTray) string {
  return ""
}

func close(m *DesktopSysTray) string {
  return ""
}
