// +build linux

package desktop

import (
	"fmt"
	"image"
)

type DesktopSysTrayAppIndicator struct {
	DesktopSysTrayGtk

	app AppIndicator
}

func DesktopSysTrayAppIndicatorNew(m *DesktopSysTray) *DesktopSysTrayAppIndicator {
	os := &DesktopSysTrayAppIndicator{*DesktopSysTrayGtkNew(m), nil}
	return os
}

func showInvoke() {
	fmt.Println("invoke")
}

func (os *DesktopSysTrayAppIndicator) show() {
	os.app = app_indicator_new("SysTrayIcon", "fallback_please", APP_INDICATOR_CATEGORY_APPLICATION_STATUS)
	go GtkMessageLoopInvoke(GSourceFunc(showInvoke))
}

func (os *DesktopSysTrayAppIndicator) hide() {
}

func (os *DesktopSysTrayAppIndicator) update() {
}

func (os *DesktopSysTrayAppIndicator) close() {
	os.close()
}

func (os *DesktopSysTrayAppIndicator) setIcon(i image.Image) {
}

