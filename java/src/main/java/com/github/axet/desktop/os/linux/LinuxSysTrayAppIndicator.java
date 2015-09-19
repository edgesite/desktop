package com.github.axet.desktop.os.linux;

import java.awt.AlphaComposite;
import java.awt.Component;
import java.awt.Graphics2D;
import java.awt.RenderingHints;
import java.awt.image.BufferedImage;
import java.util.Collections;

import javax.swing.AbstractButton;
import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JCheckBoxMenuItem;
import javax.swing.JMenu;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.Utils;
import com.github.axet.desktop.DesktopSysTray.Listener;
import com.github.axet.desktop.os.linux.handle.AppIndicator;
import com.github.axet.desktop.os.linux.handle.AppIndicatorClassStruct;
import com.github.axet.desktop.os.linux.handle.AppIndicatorInstanceStruct;
import com.github.axet.desktop.os.linux.handle.Fallback;
import com.github.axet.desktop.os.linux.handle.GBytes;
import com.github.axet.desktop.os.linux.handle.GIcon;
import com.github.axet.desktop.os.linux.handle.GMainLoop;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.github.axet.desktop.os.linux.handle.SignalCallback;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator.AppIndicatorCategory;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator.AppIndicatorStatus;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator.GtkIconSize;
import com.sun.jna.Function;
import com.sun.jna.Pointer;

/**
 * System Tray Protocol Specification
 * 
 * http://standards.freedesktop.org/systemtray-spec/systemtray-spec-latest.html
 * 
 * TODO rewrite plugin for native menus
 * 
 * see for example XSystemTrayPeer.java
 * 
 */

public class LinuxSysTrayAppIndicator extends DesktopSysTray {
    GtkWidget gtkmenu;

    JPopupMenu menu;
    String title;

    Icon icon;

    AppIndicator appindicator;
    GtkStatusIcon gicon;

    public static final Icon SpaceIcon = new ImageIcon(new BufferedImage(1, 1, BufferedImage.TYPE_INT_ARGB));

    static Thread MessageLoop = new Thread(new Runnable() {
        @Override
        public void run() {
            GMainLoop mainloop = LibAppIndicator.INSTANCE.g_main_loop_new(null, false);
            LibAppIndicator.INSTANCE.g_main_loop_run(mainloop);
        }
    });

    static {
        LibAppIndicator.INSTANCE.gtk_init(null, null);
        MessageLoop.start();
    }

    public static GIcon convertMenuImage(Icon icon) {
        BufferedImage img = Utils.createBitmap(icon);

        int menubarHeigh = 64;

        BufferedImage scaledImage = new BufferedImage(menubarHeigh, menubarHeigh, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = scaledImage.createGraphics();
        g.setRenderingHint(RenderingHints.KEY_INTERPOLATION, RenderingHints.VALUE_INTERPOLATION_BILINEAR);
        g.setRenderingHint(RenderingHints.KEY_RENDERING, RenderingHints.VALUE_RENDER_QUALITY);
        g.setRenderingHint(RenderingHints.KEY_ANTIALIASING, RenderingHints.VALUE_ANTIALIAS_ON);
        g.setComposite(AlphaComposite.getInstance(AlphaComposite.SRC_OVER));
        g.drawImage(img, 0, 0, menubarHeigh, menubarHeigh, null);
        g.dispose();

        byte[] buf = Utils.BufferedImage2Bytes(scaledImage);
        GBytes bg = LibAppIndicator.INSTANCE.g_bytes_new(buf, buf.length);
        GIcon gg = LibAppIndicator.INSTANCE.g_bytes_icon_new(bg);
        return gg;
    }

    public LinuxSysTrayAppIndicator() {
    }

    @Override
    public void setIcon(Icon icon) {
        this.icon = icon;
    }

    @Override
    public void setTitle(String title) {
        this.title = title;
    }

    @Override
    public void show() {
        updateMenus();

        if (appindicator == null) {
            appindicator = LibAppIndicator.INSTANCE.app_indicator_new(LinuxSysTrayAppIndicator.class.getSimpleName(),
                    "", AppIndicatorCategory.APP_INDICATOR_CATEGORY_APPLICATION_STATUS);

            // hacking took from https://github.com/dorkbox/SystemTray
            // we should not do this. but we can't avoid it. so lets do it :)
            AppIndicatorClassStruct aiclass = new AppIndicatorClassStruct(new AppIndicatorInstanceStruct(
                    appindicator.getPointer()).parent.g_type_instance.g_class);
            aiclass.fallback = new Fallback() {
                @Override
                public GtkStatusIcon fallback(Pointer app) {
                    gicon = LibAppIndicator.INSTANCE.gtk_status_icon_new_from_gicon(convertMenuImage(icon));
                    LibAppIndicator.INSTANCE.gtk_status_icon_set_visible(gicon, true);

                    LibAppIndicator.INSTANCE.g_signal_connect_data(gicon, "activate", new SignalCallback() {
                        @Override
                        public void signal(Pointer data) {
                            for (Listener l : Collections.synchronizedCollection(listeners)) {
                                l.mouseLeftClick();
                            }
                        }
                    }, null, null, 0);

                    LibAppIndicator.INSTANCE.g_signal_connect_data(gicon, "popup-menu", new SignalCallback() {
                        @Override
                        public void signal(Pointer data) {
                            LibAppIndicator.INSTANCE.gtk_menu_popup(gtkmenu, null, null,
                                    LibAppIndicator.gtk_status_icon_position_menu, data, 1,
                                    LibAppIndicator.INSTANCE.gtk_get_current_event_time());
                        }
                    }, null, null, 0);

                    return gicon;
                }
            };
            aiclass.write();
            LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);
        }

        LibAppIndicator.INSTANCE.app_indicator_set_status(appindicator, AppIndicatorStatus.APP_INDICATOR_STATUS_ACTIVE);
    }

    @Override
    public void update() {
        updateMenus();
        LibAppIndicator.INSTANCE.app_indicator_set_menu(appindicator, gtkmenu);

        LibAppIndicator.INSTANCE.gtk_status_icon_set_from_gicon(gicon, convertMenuImage(icon));
        LibAppIndicator.INSTANCE.g_signal_connect_data(gicon, "activate", new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                for (Listener l : Collections.synchronizedCollection(listeners)) {
                    l.mouseLeftClick();
                }
            }
        }, null, null, 0);

        LibAppIndicator.INSTANCE.g_signal_connect_data(gicon, "popup-menu", new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                LibAppIndicator.INSTANCE.gtk_menu_popup(gtkmenu, null, null,
                        LibAppIndicator.gtk_status_icon_position_menu, data, 1,
                        LibAppIndicator.INSTANCE.gtk_get_current_event_time());
            }
        }, null, null, 0);
    }

    GtkWidget createMenuItem(String n, final AbstractButton b, Boolean check, Icon img) {
        int spacing = 6;

        GtkWidget box = LibAppIndicator.INSTANCE.gtk_hbox_new(false, spacing);

        GtkWidget wicon = null;
        if (img != null) {
            wicon = LibAppIndicator.INSTANCE.gtk_image_new_from_gicon(convertMenuImage(img),
                    GtkIconSize.GTK_ICON_SIZE_MENU);
        } else {
            wicon = LibAppIndicator.INSTANCE.gtk_image_new_from_gicon(convertMenuImage(SpaceIcon),
                    GtkIconSize.GTK_ICON_SIZE_MENU);
        }
        LibAppIndicator.INSTANCE.gtk_box_pack_start(box, wicon, false, false, spacing);

        GtkWidget label = LibAppIndicator.INSTANCE.gtk_label_new(n);
        GtkWidget menu = null;

        if (check != null) {
            menu = LibAppIndicator.INSTANCE.gtk_check_menu_item_new();
            LibAppIndicator.INSTANCE.gtk_check_menu_item_set_active(menu, check.booleanValue());
        } else {
            menu = LibAppIndicator.INSTANCE.gtk_menu_item_new();
        }

        LibAppIndicator.INSTANCE.gtk_box_pack_start(box, label, false, false, spacing);
        LibAppIndicator.INSTANCE.gtk_container_add(menu, box);
        LibAppIndicator.INSTANCE.gtk_widget_show_all(menu);

        if (b != null) {
            LibAppIndicator.INSTANCE.g_signal_connect_data(menu, "activate", new SignalCallback() {
                @Override
                public void signal(Pointer data) {
                    b.doClick();
                }
            }, null, null, 0);
        }

        return menu;
    }

    void updateMenus() {
        gtkmenu = LibAppIndicator.INSTANCE.gtk_menu_new();

        for (int i = 0; i < menu.getComponentCount(); i++) {
            Component e = menu.getComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;

                GtkWidget ss = createSubmenu(sub);
                GtkWidget item1 = createMenuItem(sub.getText(), null, null, sub.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_item_set_submenu(item1, ss);
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                GtkWidget item1 = createMenuItem(ch.getText(), ch, ch.getState(), ch.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                GtkWidget item1 = createMenuItem(mi.getText(), mi, null, mi.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            }

            if (e instanceof JPopupMenu.Separator) {
                GtkWidget item1 = LibAppIndicator.INSTANCE.gtk_separator_menu_item_new();
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            }
        }
    }

    GtkWidget createSubmenu(JMenu menu) {
        GtkWidget gmenu = LibAppIndicator.INSTANCE.gtk_menu_new();

        for (int i = 0; i < menu.getMenuComponentCount(); i++) {
            Component e = menu.getMenuComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;

                GtkWidget ss = createSubmenu(sub);
                GtkWidget item = createMenuItem(sub.getText(), null, null, sub.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_item_set_submenu(item, ss);
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gmenu, item);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                GtkWidget item = createMenuItem(ch.getText(), ch, ch.getState(), ch.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gmenu, item);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                GtkWidget item = createMenuItem(mi.getText(), mi, null, mi.getIcon());
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gmenu, item);
            }

            if (e instanceof JPopupMenu.Separator) {
                GtkWidget item = LibAppIndicator.INSTANCE.gtk_separator_menu_item_new();
                LibAppIndicator.INSTANCE.gtk_menu_shell_append(gmenu, item);
            }
        }

        return gmenu;

    }

    @Override
    public void hide() {
    }

    @Override
    public void setMenu(JPopupMenu menu) {
        this.menu = menu;
    }

    @Override
    public void close() {
        hide();
    }

}
