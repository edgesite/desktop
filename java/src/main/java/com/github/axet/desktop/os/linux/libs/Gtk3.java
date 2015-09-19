package com.github.axet.desktop.os.linux.libs;

import com.github.axet.desktop.os.linux.handle.GBytes;
import com.github.axet.desktop.os.linux.handle.GIcon;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.sun.jna.Callback;
import com.sun.jna.Function;
import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Pointer;

public interface Gtk3 extends Library {

    static Gtk3 INSTANCE = (Gtk3) Native.loadLibrary("gtk-3", Gtk3.class);

    void g_signal_connect_data(Pointer item, String action, Callback callback, Pointer data, Pointer pzero1, int pzero2);

    void gtk_init(int argc, Pointer p);

    GtkWidget gtk_menu_new();

    void gtk_menu_shell_append(GtkWidget menu, GtkWidget item);

    GtkWidget gtk_menu_item_new_with_label(String s);

    GtkWidget gtk_check_menu_item_new_with_label(String s);

    String gtk_menu_item_get_label(GtkWidget item);

    void gtk_widget_show(GtkWidget item);

    Pointer g_main_loop_new(Pointer context, boolean is_running);

    void g_main_loop_run(Pointer loop);

    GtkStatusIcon gtk_status_icon_new_from_gicon(GIcon icon);

    GBytes g_bytes_new(byte[] buf, int size);

    GIcon g_bytes_icon_new(GBytes bytes);

    void gtk_menu_popup(GtkWidget m, GtkWidget parent, GtkWidget parentitem, Function func, Pointer data, int button,
            int time);

    int gtk_get_current_event_time();

    void gtk_status_icon_set_visible(GtkStatusIcon icon, boolean b);
}
