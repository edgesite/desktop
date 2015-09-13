// +build darwin

package desktop

/*
extern void DesktopSysTrayActionOSXActionMap(void*, void*);
*/
import "C"

import (
	"unsafe"
)

//
// register
//

var DesktopSysTrayActionOSXClassReg unsafe.Pointer = Runtime_objc_allocateClassPair(NSObjectClass, "DesktopSysTrayActionOSXClass", 0)
var DesktopSysTrayActionOSXActionReg unsafe.Pointer = Runtime_sel_registerName("action")

var DesktopSysTrayActionOSXMap = make(map[unsafe.Pointer]*DesktopSysTrayActionOSX)

//export DesktopSysTrayActionOSXActionMap
func DesktopSysTrayActionOSXActionMap(id unsafe.Pointer, sel unsafe.Pointer) {
	if sel == DesktopSysTrayActionOSXActionReg {
		DesktopSysTrayActionOSXMap[id].Action()
	}
}

func DesktopSysTrayActionOSXRegister() bool {
	if !Runtime_class_addMethod(DesktopSysTrayActionOSXClassReg, DesktopSysTrayActionOSXActionReg, C.DesktopSysTrayActionOSXActionMap, "v@:") {
		panic("problem initalizing class")
	}
	Runtime_objc_registerClassPair(DesktopSysTrayActionOSXClassReg)

	return true
}

var DesktopSysTrayActionOSXRegistred bool = DesktopSysTrayActionOSXRegister()

//
// object
//

var DesktopSysTrayActionOSXClass unsafe.Pointer = Runtime_objc_lookUpClass("DesktopSysTrayActionOSXClass")
var DesktopSysTrayActionOSXAction unsafe.Pointer = Runtime_sel_getUid("action")

type DesktopSysTrayActionOSX struct {
	NSObject

	Menu *Menu
}

func DesktopSysTrayActionOSXNew(mn *Menu) *DesktopSysTrayActionOSX {
	m := DesktopSysTrayActionOSXPointer(Runtime_class_createInstance(DesktopSysTrayActionOSXClass, 0))

	m.Menu = mn

	return m
}

func DesktopSysTrayActionOSXPointer(p unsafe.Pointer) *DesktopSysTrayActionOSX {
	m := &DesktopSysTrayActionOSX{NSObjectPointer(p), nil}

	DesktopSysTrayActionOSXMap[m.Pointer] = m

	return m
}

func (m *DesktopSysTrayActionOSX) Action() {
	m.Menu.Action(m.Menu)
}

func (m *DesktopSysTrayActionOSX) Release() {
	delete(DesktopSysTrayActionOSXMap, m.Pointer)
	m.NSObject.Release()
}
