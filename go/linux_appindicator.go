// +build linux

package desktop

/*
#include <stdlib.h>

#include "appindicator.h"

#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
*/
import "C"

import (
	"unsafe"
)

const (
        APP_INDICATOR_CATEGORY_APPLICATION_STATUS = 0;
        APP_INDICATOR_CATEGORY_COMMUNICATIONS = 1;
        APP_INDICATOR_CATEGORY_SYSTEM_SERVICES = 2;
        APP_INDICATOR_CATEGORY_HARDWARE = 3;
        APP_INDICATOR_CATEGORY_OTHER = 4;

        APP_INDICATOR_STATUS_PASSIVE = 0;
        APP_INDICATOR_STATUS_ACTIVE = 1;
        APP_INDICATOR_STATUS_ATTENTION = 2;
)

func app_indicator_new(id string, icon_name string, category int) uintptr {
	n := C.CString(id)
	defer C.free(unsafe.Pointer(n))
	k := C.CString(icon_name)
	defer C.free(unsafe.Pointer(k))
	return uintptr(C.app_indicator_new(n, k, (C.int)(category)))
}

func app_indicator_set_icon_theme_path(app uintptr, path string) {
	n := C.CString(path)
	defer C.free(unsafe.Pointer(n))
	C.app_indicator_set_icon_theme_path(unsafe.Pointer(app), n)
}

func app_indicator_set_menu(app uintptr, menu uintptr) {
	C.app_indicator_set_menu(unsafe.Pointer(app), unsafe.Pointer(menu))
}

func app_indicator_set_icon_full(app uintptr, name string, desc string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	k := C.CString(desc)
	defer C.free(unsafe.Pointer(k))
	C.app_indicator_set_icon_full(unsafe.Pointer(app), n, k)
}

func app_indicator_set_title(app uintptr, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.app_indicator_set_title(unsafe.Pointer(app), n)
}

func app_indicator_set_status(app uintptr, status int){
	C.app_indicator_set_status(unsafe.Pointer(app), (C.int)(status))
}

