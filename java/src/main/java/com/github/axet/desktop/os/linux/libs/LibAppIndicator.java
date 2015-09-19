package com.github.axet.desktop.os.linux.libs;

import com.github.axet.desktop.os.linux.handle.AppIndicator;
import com.github.axet.desktop.os.linux.handle.GBytes;
import com.github.axet.desktop.os.linux.handle.GIcon;
import com.github.axet.desktop.os.linux.handle.GMainLoop;
import com.github.axet.desktop.os.linux.handle.GObject;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.sun.jna.Callback;
import com.sun.jna.Function;
import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Pointer;

public interface LibAppIndicator extends Library {

    public static String NAME = "appindicator";

    public static LibAppIndicator INSTANCE = (LibAppIndicator) Native.loadLibrary(NAME, LibAppIndicator.class);

    public static Function gtk_status_icon_position_menu = Function.getFunction(LibAppIndicator.NAME,
            "gtk_status_icon_position_menu");

    public interface AppIndicatorCategory {
        public static final int APP_INDICATOR_CATEGORY_APPLICATION_STATUS = 0;

        public static final int APP_INDICATOR_CATEGORY_COMMUNICATIONS = 1;

        public static final int APP_INDICATOR_CATEGORY_SYSTEM_SERVICES = 2;

        public static final int APP_INDICATOR_CATEGORY_HARDWARE = 3;

        public static final int APP_INDICATOR_CATEGORY_OTHER = 4;
    }

    public interface AppIndicatorStatus {
        public static final int APP_INDICATOR_STATUS_PASSIVE = 0;
        public static final int APP_INDICATOR_STATUS_ACTIVE = 1;
        public static final int APP_INDICATOR_STATUS_ATTENTION = 2;
    }

    public interface GtkOrientation {
        public static final int GTK_ORIENTATION_HORIZONTAL = 0;
        public static final int GTK_ORIENTATION_VERTICAL = 1;
    }

    public interface GtkIconSize {
        // Invalid size.
        public static final int GTK_ICON_SIZE_INVALID = 0;

        // Size appropriate for menus (16px).
        public static final int GTK_ICON_SIZE_MENU = 1;

        // Size appropriate for small toolbars (16px).
        public static final int GTK_ICON_SIZE_SMALL_TOOLBAR = 2;

        // Size appropriate for large toolbars (24px)
        public static final int GTK_ICON_SIZE_LARGE_TOOLBAR = 3;

        // Size appropriate for buttons (16px)
        public static final int GTK_ICON_SIZE_BUTTON = 4;

        // Size appropriate for drag and drop (32px)
        public static final int GTK_ICON_SIZE_DND = 5;

        // Size appropriate for dialogs (48px)
        public static final int GTK_ICON_SIZE_DIALOG = 6;
    }

    AppIndicator app_indicator_new(String id, String icon_name, int cat);

    void app_indicator_set_icon_theme_path(AppIndicator app, String path);

    void app_indicator_set_menu(AppIndicator app, GtkWidget menu);

    void app_indicator_set_icon_full(AppIndicator app, String name, String desc);

    void app_indicator_set_title(AppIndicator app, String title);

    void app_indicator_set_status(AppIndicator app, int status);

    //
    // gtk calls
    //

    void g_signal_connect_data(GObject item, String action, Callback callback, Pointer data, Pointer pzero1, int pzero2);

    void gtk_init(Pointer pargc, Pointer pargv);

    void g_object_ref(GObject p);

    void g_object_unref(GObject p);

    int gtk_get_current_event_time();

    // menus

    GtkWidget gtk_menu_new();

    void gtk_menu_shell_append(GtkWidget menu, GtkWidget item);

    GtkWidget gtk_separator_menu_item_new();

    GtkWidget gtk_menu_item_new();

    GtkWidget gtk_menu_item_new_with_label(String s);

    GtkWidget gtk_check_menu_item_new_with_label(String s);

    String gtk_menu_item_get_label(GtkWidget item);

    void gtk_menu_item_set_submenu(GtkWidget menu, GtkWidget item);

    void gtk_menu_popup(GtkWidget m, GtkWidget parent, GtkWidget parentitem, Function func, Pointer data, int button,
            int time);

    void gtk_widget_show(GtkWidget item);

    GtkWidget gtk_hbox_new(boolean homogeneous, int spacing);

    void gtk_box_pack_start(GtkWidget box, GtkWidget item, boolean expand, boolean fill, int padding);

    void gtk_box_pack_end(GtkWidget box, GtkWidget item, boolean expand, boolean fill, int padding);

    GtkWidget gtk_label_new(String s);

    void gtk_label_set_text(GtkWidget label, String s);

    String gtk_label_get_text(GtkWidget label);

    void gtk_container_add(GtkWidget container, GtkWidget widget);

    void gtk_widget_show_all(GtkWidget container);

    GtkWidget gtk_check_menu_item_new();

    void gtk_check_menu_item_set_active(GtkWidget menu, boolean b);

    // status icon

    GtkStatusIcon gtk_status_icon_new_from_gicon(GIcon icon);

    void gtk_status_icon_set_from_gicon(GtkStatusIcon s, GIcon i);

    void gtk_status_icon_set_visible(GtkStatusIcon icon, boolean b);

    GtkWidget gtk_image_new();

    GtkWidget gtk_image_new_from_gicon(GIcon g, int size);

    // GBytes

    GBytes g_bytes_new(byte[] buf, int size);

    GIcon g_bytes_icon_new(GBytes bytes);

    void g_bytes_unref(GBytes b);

    // loop

    GMainLoop g_main_loop_new(Pointer context, boolean is_running);

    void g_main_loop_run(GMainLoop loop);

    void g_main_loop_quit(GMainLoop loop);
}
