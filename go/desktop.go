package desktop

import (
	"image"
)

//
// Desktop Folders
//

// Config folder
//
//  - osx: /Users/user/Library/Application Support
//  - windows: C:\Users\user\AppData\Local
//  - linux: /home/user/.config
func GetAppDataFolder() string {
	return getAppDataFolder()
}

// Home folder
//
//  - osx: /Users/user
//  - windows: C:\Users\user
//  - linux: /home/user
func GetHomeFolder() string {
	return getHomeFolder()
}

// Documents folder
//
//  - osx: /Users/user/Documents
//  - windows: C:\Users\user\Documents
//  - linux: /home/user/Documents
func GetDocumentsFolder() string {
	return getDocumentsFolder()
}

// Downloads folder
//
//  - osx: /Users/user/Downloads
//  - windows: C:\Users\user\Downloads
//  - linux: /home/user/Desktop
func GetDownloadsFolder() string {
	return getDownloadsFolder()
}

// Desktop folder
//
//  - osx: /Users/user/Desktop
//  - windows: C:\Users\user\Desktop
//  - linux: /home/user/Desktop
func GetDesktopFolder() string {
	return getDesktopFolder()
}

//
// Main function
//
// Need to keep messages loop running. Have to be run on main thread.
// All GUI applications need that. So if you plan to use SysTray call
// this function from main function.
//

func Main() {
	desktopMain()
}

//
// SysTrayIcon or NSStatusBar or Notification Area
//

const (
	MenuItem      = 1
	MenuSeparator = 2
	MenuCheckBox  = 3
)

type MenuAction func(*Menu)

type Menu struct {
	Menu    []Menu
	Action  MenuAction
	State   bool
	Type    int
	Enabled bool
	Name    string
	Icon    image.Image
}

type DesktopSysTray struct {
	Listeners map[DesktopSysTrayListener]bool
	Icon      []byte
	Title     string
	Menu      []Menu

	// os specific structs
	os interface{}
}

type DesktopSysTrayListener interface {
	MouseLeftClick()

	MouseLeftDoubleClick()

	// We do not handle right clicks, because:
	//
	// 1) Icon is binded to context menu anyway.
	//
	// 2) On Windows if you call showContextMenu from java thread, HMENU bugged
	// and you can't use it.
	//
	// 3) Mac OSX does not support showing context menu programmatically.
}

func DesktopSysTrayNew() *DesktopSysTray {
	return desktopSysTrayNew()
}

func (m *DesktopSysTray) AddListener(l DesktopSysTrayListener) {
	m.Listeners[l] = true
}

func (m *DesktopSysTray) RemoveListener(l DesktopSysTrayListener) {
	delete(m.Listeners, l)
}

func (m *DesktopSysTray) SetIcon(icon image.Image) {
	setIcon(m, icon)
}

func (m *DesktopSysTray) SetTitle(title string) {
	m.Title = title
}

func (m *DesktopSysTray) Show() {
	show(m)
}

func (m *DesktopSysTray) Update() {
	update(m)
}

func (m *DesktopSysTray) Hide() {
	hide(m)
}

func (m *DesktopSysTray) SetMenu(menu []Menu) {
	m.Menu = menu
}

func (m *DesktopSysTray) Close() {
	close(m)
}
