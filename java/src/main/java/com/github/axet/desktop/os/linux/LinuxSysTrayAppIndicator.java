package com.github.axet.desktop.os.linux;

import com.github.axet.desktop.os.linux.handle.AppIndicator;
import com.github.axet.desktop.os.linux.handle.AppIndicatorClassStruct;
import com.github.axet.desktop.os.linux.handle.AppIndicatorInstanceStruct;
import com.github.axet.desktop.os.linux.handle.Fallback;
import com.github.axet.desktop.os.linux.handle.GtkIconSet;
import com.github.axet.desktop.os.linux.handle.GtkMessageLoop;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator.AppIndicatorCategory;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator.AppIndicatorStatus;
import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;

public class LinuxSysTrayAppIndicator extends LinuxSysTrayGtk {

    AppIndicator appindicator;
    GtkIconSet iconset;

    void createAppIndicator() {
        if (appindicator == null) {
            appindicator = LibAppIndicator.INSTANCE.app_indicator_new(LinuxSysTrayAppIndicator.class.getSimpleName(),
                    "fallback_please", AppIndicatorCategory.APP_INDICATOR_CATEGORY_APPLICATION_STATUS);

            // hacking took from https://github.com/dorkbox/SystemTray
            // we should not do this. but we can't avoid it. so lets do
            // it :)
            AppIndicatorClassStruct aiclass = new AppIndicatorClassStruct(new AppIndicatorInstanceStruct(
                    appindicator.getPointer()).parent.g_type_instance.g_class);
            aiclass.fallback = new Fallback() {
                @Override
                public GtkStatusIcon fallback(Pointer app) {
                    gtkstatusicon = createGStatusIcon();
                    LibGtk.INSTANCE.gtk_status_icon_set_visible(gtkstatusicon, true);
                    return gtkstatusicon;
                }
            };
            aiclass.write();
            LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);
        }
    }

    void updateIcon() {
        if (icon == null) {
            return;
        }

        if (iconset == null)
            iconset = new GtkIconSet();
        String p = iconset.addIcon(icon);
        LibAppIndicator.INSTANCE.app_indicator_set_icon_theme_path(appindicator, iconset.getPath());
        LibAppIndicator.INSTANCE.app_indicator_set_icon_full(appindicator, p, getClass().getSimpleName());
    }

    //
    // public
    //

    public LinuxSysTrayAppIndicator() {
    }

    @Override
    public void show() {
        GtkMessageLoop.invokeLater(new Runnable() {
            @Override
            public void run() {
                updateMenus();

                createAppIndicator();

                LibAppIndicator.INSTANCE.app_indicator_set_status(appindicator,
                        AppIndicatorStatus.APP_INDICATOR_STATUS_ACTIVE);
                LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);

                updateIcon();
            }
        });
    }

    @Override
    public void update() {
        GtkMessageLoop.invokeLater(new Runnable() {
            @Override
            public void run() {
                updateMenus();

                LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);

                if (gtkstatusicon != null) {
                    LibGtk.INSTANCE.gtk_status_icon_set_from_gicon(gtkstatusicon, convertMenuImage(icon));
                    LibGtk.INSTANCE.gtk_status_icon_set_tooltip_text(gtkstatusicon, title);
                }

                updateIcon();
            }
        });
    }

    @Override
    public void hide() {
        GtkMessageLoop.invokeLater(new Runnable() {
            @Override
            public void run() {
                if (gtkstatusicon != null) {
                    LibGtk.INSTANCE.gtk_status_icon_set_visible(gtkstatusicon, false);
                }

                LibAppIndicator.INSTANCE.app_indicator_set_status(appindicator,
                        AppIndicatorStatus.APP_INDICATOR_STATUS_PASSIVE);
            }
        });
    }

}
