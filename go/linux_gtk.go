// +build linux

package desktop

/*
#include <stdlib.h>
#include <stdio.h>

#include "gtk.h"

extern void* gsourcefunc(void*);
extern void signal_activate(void*, void*);
extern void signal_popup_menu(void*, void*, void*,void*);

#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
*/
import "C"

import (
	"unsafe"
)

const (
	GTK_ORIENTATION_HORIZONTAL  = 0
	GTK_ORIENTATION_VERTICAL    = 1
	GTK_ICON_SIZE_INVALID       = 0
	GTK_ICON_SIZE_MENU          = 1
	GTK_ICON_SIZE_SMALL_TOOLBAR = 2
	GTK_ICON_SIZE_LARGE_TOOLBAR = 3
	GTK_ICON_SIZE_BUTTON        = 4
	GTK_ICON_SIZE_DND           = 5
	GTK_ICON_SIZE_DIALOG        = 6
)

type GObject unsafe.Pointer
type GtkWidget unsafe.Pointer
type GMainLoop unsafe.Pointer
type GMainContext unsafe.Pointer
type GSourceFunc func()
type GPointer unsafe.Pointer
type GtkStatusIcon unsafe.Pointer
type GIcon unsafe.Pointer
type GBytes unsafe.Pointer

func gtk_init() {
	C.gtk_init(0, 0)
}

//export gsourcefunc
func gsourcefunc(p unsafe.Pointer) unsafe.Pointer {
	f := *(*GSourceFunc)(unsafe.Pointer(p))
	f()
	return Arg(false)
}

//export signal_activate
func signal_activate(p, p1 unsafe.Pointer) {
	f := *(*GSourceFunc)(unsafe.Pointer(p1))
	f()
}

//export signal_popup_menu
func signal_popup_menu(p, p1, p2, p3 unsafe.Pointer) {
	f := *(*GSourceFunc)(unsafe.Pointer(p3))
	f()
}

func g_signal_connect_activate(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("activate")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_activate), Arg(fn), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_connect_popup(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("popup-menu")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_popup_menu), Arg(fn), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_emit_by_name(item GtkWidget, action string) {
	n := C.CString(action)
	defer C.free(unsafe.Pointer(n))
	C.g_signal_emit_by_name(Arg(item), n)
}

func g_object_ref(p GObject) {
	C.g_object_ref(Arg(p))
}

func g_object_unref(p GObject) {
	C.g_object_unref(Arg(p))
}

func gtk_widget_destroy(p GtkWidget) {
	C.gtk_widget_destroy(Arg(p))
}

func gtk_get_current_event_time() int {
	return int(C.gtk_get_current_event_time())
}

func gtk_menu_new() GtkWidget {
	return GtkWidget(C.gtk_menu_new())
}

func gtk_menu_shell_append(menu GtkWidget, item GtkWidget) {
	C.gtk_menu_shell_append(Arg(menu), Arg(item))
}

func gtk_separator_menu_item_new() GtkWidget {
	return GtkWidget(C.gtk_separator_menu_item_new())
}

func gtk_menu_item_new() GtkWidget {
	return GtkWidget(C.gtk_menu_item_new())
}

func gtk_menu_item_new_with_label(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_menu_item_new_with_label(n))
}

func gtk_check_menu_item_new_with_label(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_check_menu_item_new_with_label(n))
}

func gtk_menu_item_get_label(item GtkWidget) string {
	return C.GoString(C.gtk_menu_item_get_label(Arg(item)))
}

func gtk_menu_item_set_submenu(menu GtkWidget, item GtkWidget) {
	C.gtk_menu_item_set_submenu(Arg(menu), Arg(item))
}

func gtk_menu_popup(m GtkWidget, parent GtkWidget, parentitem GtkWidget, fn GPointer, data GPointer, button int, time int) {
	C.gtk_menu_popup(Arg(m), Arg(parent), Arg(parentitem), Arg(fn), Arg(data), C.int(button), C.int(time))
}

var gtk_status_icon_position_menu = GPointer(C.gtk_status_icon_position_menu)

func gtk_widget_show(item GtkWidget) {
	C.gtk_widget_show(Arg(item))
}

func gtk_hbox_new(homogeneous bool, spacing int) GtkWidget {
	return GtkWidget(C.gtk_hbox_new(C.bool(Bool2Int[homogeneous]), C.int(spacing)))
}

func gtk_box_pack_start(box GtkWidget, item GtkWidget, expand bool, fill bool, padding int) {
	C.gtk_box_pack_start(Arg(box), Arg(item), C.bool(Bool2Int[expand]), C.bool(Bool2Int[fill]), C.int(padding))
}

func gtk_box_pack_end(box GtkWidget, item GtkWidget, expand bool, fill bool, padding int) {
	C.gtk_box_pack_end(Arg(box), Arg(item), C.bool(Bool2Int[expand]), C.bool(Bool2Int[fill]), C.int(padding))
}

func gtk_label_new(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_label_new(n))
}

func gtk_label_set_text(label GtkWidget, s string) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	C.gtk_label_set_text(Arg(label), n)
}

func gtk_label_get_text(label GtkWidget) string {
	return C.GoString(C.gtk_label_get_text(Arg(label)))
}

func gtk_container_add(container GtkWidget, widget GtkWidget) {
	C.gtk_container_add(Arg(container), Arg(widget))
}

func gtk_widget_show_all(container GtkWidget) {
	C.gtk_widget_show_all(Arg(container))
}

func gtk_check_menu_item_new() GtkWidget {
	return GtkWidget(GtkWidget(C.gtk_check_menu_item_new()))
}

func gtk_check_menu_item_set_active(menu GtkWidget, b bool) {
	C.gtk_check_menu_item_set_active(Arg(menu), C.bool(Bool2Int[b]))
}

func gtk_widget_set_sensitive(item GtkWidget, b bool) {
	C.gtk_widget_set_sensitive(Arg(item), C.bool(Bool2Int[b]))
}

func gtk_status_icon_new_from_gicon(icon GIcon) GtkWidget {
	return GtkWidget(C.gtk_status_icon_new_from_gicon(Arg(icon)))
}

func gtk_status_icon_set_from_gicon(s GtkWidget, i GIcon) {
	C.gtk_status_icon_set_from_gicon(Arg(s), Arg(i))
}

func gtk_status_icon_set_visible(icon GtkWidget, b bool) {
	C.gtk_status_icon_set_visible(Arg(icon), C.bool(Bool2Int[b]))
}

func gtk_image_new() GtkWidget {
	return GtkWidget(C.gtk_image_new())
}

func gtk_image_new_from_gicon(g GIcon, size int) GtkWidget {
	return GtkWidget(C.gtk_image_new_from_gicon(Arg(g), C.int(size)))
}

func gtk_status_icon_set_title(icon GtkWidget, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.gtk_status_icon_set_title(Arg(icon), n)
}

func gtk_status_icon_get_title(icon GtkWidget) string {
	return C.GoString(C.gtk_status_icon_get_title(Arg(icon)))
}

func gtk_status_icon_set_tooltip_text(icon GtkWidget, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.gtk_status_icon_set_tooltip_text(Arg(icon), n)
}

func gtk_status_icon_get_tooltip_text(icon GtkWidget) string {
	return C.GoString(C.gtk_status_icon_get_tooltip_text(Arg(icon)))
}

func g_bytes_new(buf []byte, size int) GBytes {
	return GBytes(C.g_bytes_new(Arg(buf), C.int(size)))
}

func g_bytes_icon_new(bytes GBytes) GIcon {
	return GIcon(C.g_bytes_icon_new(Arg(bytes)))
}

func g_bytes_unref(b GBytes) {
	C.g_bytes_unref(Arg(b))
}

func g_main_loop_new(context GMainContext, is_running bool) GMainLoop {
	return GMainLoop(C.g_main_loop_new(Arg(context), Arg(Bool2Int[is_running])))
}

func g_main_loop_run(loop GMainLoop) {
	C.g_main_loop_run(Arg(loop))
}

func g_main_loop_quit(loop GMainLoop) {
	C.g_main_loop_quit(Arg(loop))
}

func g_main_loop_get_context(loop GMainLoop) GMainContext {
	return GMainContext(C.g_main_loop_get_context(Arg(loop)))
}

func g_main_context_invoke(c GMainContext, fn *GSourceFunc) {
	C.g_main_context_invoke(Arg(c), Arg(C.gsourcefunc), Arg(fn))
}
