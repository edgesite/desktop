// +build linux

package desktop

import (
	"image"
	"fmt"
)

type DesktopSysTrayAppIndicator struct {
	DesktopSysTrayGtk
}

func (os *DesktopSysTrayAppIndicator) show() {
	fmt.Println("start");
	gtk_init()
	app_indicator_new("abc", "abc", 0)
	fmt.Println("done");
}

func (os *DesktopSysTrayAppIndicator) hide() {
}

func (os *DesktopSysTrayAppIndicator) update() {
}

func (os *DesktopSysTrayAppIndicator) close() {
}

func (os *DesktopSysTrayAppIndicator) setIcon(i image.Image) {
}

