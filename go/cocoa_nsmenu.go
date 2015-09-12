// +build darwin

package desktop

import "C"

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ApplicationKit/Classes/NSImage_Class

var NSMenuClass unsafe.Pointer = Runtime_objc_lookUpClass("NSMenu")
var NSMenuAddItem unsafe.Pointer = Runtime_sel_getUid("addItem:")
var NSMenuSetAutoenablesItems unsafe.Pointer = Runtime_sel_getUid("setAutoenablesItems:")
var NSMenuItemAtIndex unsafe.Pointer = Runtime_sel_getUid("itemAtIndex:")
var NSMenuInsertItemAtIndex unsafe.Pointer = Runtime_sel_getUid("insertItem:atIndex:")

type NSMenu struct {
	NSObject
}

func NSMenuNew() NSMenu {
  return NSMenu{NSObjectPointer(Runtime_class_createInstance(NSMenuClass, 0))}
}

func NSMenuPointer(p unsafe.Pointer) NSMenu {
	return NSMenu{NSObjectPointer(p)}
}

func (m *NSMenu) NSMenuAddItem(i NSMenuItem) {
  Runtime_objc_msgSend(m.Pointer, i.Pointer)
}

func (m *NSMenu) NSMenuInsertItemAtIndex(i NSMenuItem, in int) {
  Runtime_objc_msgSend(m.Pointer, i.Pointer, Int2Pointer(in))
}

func (m *NSMenu) NSMenuNSMenuSetAutoenablesItems(b bool) {
  Runtime_objc_msgSend(m.Pointer, Bool2Pointer(b))
}

func (m *NSMenu) NSMenuItemAtIndex(i int) NSMenuItem {
  return NSMenuItemPointer(Runtime_objc_msgSend(m.Pointer, Int2Pointer(i)))
}
