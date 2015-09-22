package com.github.axet.desktop.os.linux;

import javax.swing.Icon;

import com.github.axet.desktop.os.linux.handle.AppIndicator;
import com.github.axet.desktop.os.linux.handle.AppIndicatorClassStruct;
import com.github.axet.desktop.os.linux.handle.AppIndicatorInstanceStruct;
import com.github.axet.desktop.os.linux.handle.Fallback;
import com.github.axet.desktop.os.linux.handle.GSourceFunc;
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
            AppIndicatorInstanceStruct inst = new AppIndicatorInstanceStruct(appindicator.getPointer());
            AppIndicatorClassStruct aiclass = new AppIndicatorClassStruct(inst.parent.g_type_instance.g_class);
            aiclass.fallback = new Fallback() {
                @Override
                public GtkStatusIcon fallback(Pointer app) {
                    gtkstatusicon = createGStatusIcon();
                    LibGtk.INSTANCE.gtk_status_icon_set_visible(gtkstatusicon, true);
                    return gtkstatusicon;
                }
            };
            aiclass.write();
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

        if (gtkstatusicon != null) {
            LibGtk.INSTANCE.gtk_status_icon_set_from_gicon(gtkstatusicon, convertMenuImage(icon));
            LibGtk.INSTANCE.gtk_status_icon_set_tooltip_text(gtkstatusicon, title);
        }
    }

    //
    // public
    //

    public LinuxSysTrayAppIndicator() {
    }

    GSourceFunc show = new GSourceFunc() {
        @Override
        public boolean invoke(Pointer data) {
            updateMenus();

            createAppIndicator();

            updateIcon();

            LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);
            LibAppIndicator.INSTANCE.app_indicator_set_status(appindicator,
                    AppIndicatorStatus.APP_INDICATOR_STATUS_ACTIVE);
            return false;
        }
    };

    @Override
    public void show() {
        GtkMessageLoop.invokeLater(show, null);
    }

    GSourceFunc update = new GSourceFunc() {
        @Override
        public boolean invoke(Pointer data) {
            updateMenus();

            LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);

            updateIcon();
            return false;
        }
    };

    @Override
    public void update() {
        GtkMessageLoop.invokeLater(update, null);
    }

    GSourceFunc hide = new GSourceFunc() {
        @Override
        public boolean invoke(Pointer data) {
            if (gtkstatusicon != null) {
                LibGtk.INSTANCE.gtk_status_icon_set_visible(gtkstatusicon, false);
            }

            LibAppIndicator.INSTANCE.app_indicator_set_status(appindicator,
                    AppIndicatorStatus.APP_INDICATOR_STATUS_PASSIVE);
            return false;
        }
    };

    @Override
    public void hide() {
        GtkMessageLoop.invokeLater(hide, null);
    }

    @Override
    public void close() {
        if (appindicator != null) {
            appindicator.unref();
            appindicator = null;
        }
        
        super.close();
    }

    GSourceFunc setIcon = new GSourceFunc() {
        @Override
        public boolean invoke(Pointer data) {
            updateIcon();

            return false;
        }
    };

    @Override
    public void setIcon(Icon icon) {
        this.icon = icon;

        GtkMessageLoop.invokeLater(setIcon, null);
    }
}
