package com.github.axet.desktop.os.linux.libs;

import com.github.axet.desktop.os.linux.handle.AppIndicator;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.sun.jna.Library;
import com.sun.jna.Native;

public interface LibAppIndicator extends Library {

    public static LibAppIndicator INSTANCE = (LibAppIndicator) Native.loadLibrary(LibGtkName.getName(),
            LibAppIndicator.class);

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

    AppIndicator app_indicator_new(String id, String icon_name, int cat);

    void app_indicator_set_icon_theme_path(AppIndicator app, String path);

    void app_indicator_set_menu(AppIndicator app, GtkWidget menu);

    void app_indicator_set_icon_full(AppIndicator app, String name, String desc);

    void app_indicator_set_title(AppIndicator app, String title);

    void app_indicator_set_status(AppIndicator app, int status);

}
