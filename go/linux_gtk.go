// +build linux

package desktop

/*
#include <stdlib.h>
#include <stdio.h>

#include "gtk.h"

extern void* gsourcefunc(void*);

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

func gtk_init() {
	C.gtk_init(0, 0)
}

//export gsourcefunc
func gsourcefunc(p unsafe.Pointer) unsafe.Pointer {
	var pp = &p
	f := *(*GSourceFunc)(unsafe.Pointer(&pp))
	f()
	return Arg(false)
}

func g_signal_connect(item unsafe.Pointer, action string, callback GSourceFunc) {
	n := C.CString(action)
	defer C.free(unsafe.Pointer(n))
	C.g_signal_connect_data(item, n, Arg(C.gsourcefunc), Arg(callback), NULL, C.int(0))
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

/*
void* gtk_menu_item_new_with_label(const char* s);
void* gtk_check_menu_item_new_with_label(const char* s);
const char* gtk_menu_item_get_label(void* item);
void gtk_menu_item_set_submenu(void* menu, void* item);
void gtk_menu_popup(void* m, void* parent, void* parentitem, void* func, void* data, int button, int time);
void gtk_widget_show(void* item);
void* gtk_hbox_new(bool homogeneous, int spacing);
void gtk_box_pack_start(void* box, void* item, bool expand, bool fill, int padding);
void gtk_box_pack_end(void* box, void* item, bool expand, bool fill, int padding);
void* gtk_label_new(const char* s);
void gtk_label_set_text(void* label, const char* s);
const char* gtk_label_get_text(void* label);
void gtk_container_add(void* container, void* widget);
void gtk_widget_show_all(void* container);
void* gtk_check_menu_item_new();
void gtk_check_menu_item_set_active(void* menu, bool b);
void gtk_widget_set_sensitive(void* item, bool b);

// status icon
void* gtk_status_icon_new_from_gicon(void* icon);
void gtk_status_icon_set_from_gicon(void* s, void* i);
void gtk_status_icon_set_visible(void* icon, bool b);
void* gtk_image_new();
void* gtk_image_new_from_gicon(void* g, int size);
void gtk_status_icon_set_title(void* icon, const char* title);
const char* gtk_status_icon_get_title(void* icon);
void gtk_status_icon_set_tooltip_text(void* icon, const char* title);
const char* gtk_status_icon_get_tooltip_text(void* icon);
// GBytes
void* g_bytes_new(void* buf, int size);
void* g_bytes_icon_new(void* bytes);
void g_bytes_unref(void* b);


// threads
void gdk_threads_init();
void gdk_threads_enter();
void gdk_threads_leave();
*/

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

func g_main_context_invoke(c GMainContext, fn GSourceFunc) {
	C.g_main_context_invoke(Arg(c), Arg(C.gsourcefunc), Arg(fn))
}

