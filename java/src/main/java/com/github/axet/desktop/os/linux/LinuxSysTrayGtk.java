package com.github.axet.desktop.os.linux;

import java.awt.AlphaComposite;
import java.awt.Component;
import java.awt.Graphics2D;
import java.awt.RenderingHints;
import java.awt.image.BufferedImage;
import java.util.ArrayList;
import java.util.Collections;

import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JCheckBoxMenuItem;
import javax.swing.JMenu;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.Utils;
import com.github.axet.desktop.os.linux.handle.GBytes;
import com.github.axet.desktop.os.linux.handle.GIcon;
import com.github.axet.desktop.os.linux.handle.GSourceFunc;
import com.github.axet.desktop.os.linux.handle.GtkMenuItem;
import com.github.axet.desktop.os.linux.handle.GtkMessageLoop;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.github.axet.desktop.os.linux.handle.SignalCallback;
import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.github.axet.desktop.os.linux.libs.LibGtk.GtkIconSize;
import com.sun.jna.Pointer;

public class LinuxSysTrayGtk extends DesktopSysTray {
    GtkWidget gtkmenu;

    JPopupMenu menu;
    String title;

    Icon icon;

    GtkStatusIcon gtkstatusicon;

    ArrayList<GtkMenuItem> menukepper = new ArrayList<GtkMenuItem>();

    static final String activate = "activate";

    static final String popup_menu = "popup-menu";

    public static final Icon SpaceIcon = new ImageIcon(new BufferedImage(1, 1, BufferedImage.TYPE_INT_ARGB));

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
        GBytes bg = LibGtk.INSTANCE.g_bytes_new(buf, buf.length);
        GIcon gg = LibGtk.INSTANCE.g_bytes_icon_new(bg);
        return gg;
    }

    GtkWidget createMenuItem(final JMenuItem item) {
        Icon img = item.getIcon();

        if (img == null) {
            img = SpaceIcon;
        }

        int spacing = 6;

        GtkWidget box = LibGtk.INSTANCE.gtk_hbox_new(false, spacing);

        GtkWidget wicon = LibGtk.INSTANCE.gtk_image_new_from_gicon(convertMenuImage(img),
                GtkIconSize.GTK_ICON_SIZE_MENU);
        LibGtk.INSTANCE.gtk_box_pack_start(box, wicon, false, false, spacing);

        GtkWidget label = LibGtk.INSTANCE.gtk_label_new(item.getText());
        GtkWidget menu = null;

        if (item instanceof JCheckBoxMenuItem) {
            menu = LibGtk.INSTANCE.gtk_check_menu_item_new();
            LibGtk.INSTANCE.gtk_check_menu_item_set_active(menu, ((JCheckBoxMenuItem) item).getState());
        } else {
            menu = LibGtk.INSTANCE.gtk_menu_item_new();
        }

        if (!item.isEnabled()) {
            LibGtk.INSTANCE.gtk_widget_set_sensitive(menu, false);
        }

        LibGtk.INSTANCE.gtk_box_pack_start(box, label, false, false, spacing);
        LibGtk.INSTANCE.gtk_container_add(menu, box);
        LibGtk.INSTANCE.gtk_widget_show_all(menu);

        GtkMenuItem gtkmenu = new GtkMenuItem(menu.getPointer());
        gtkmenu.activate = new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                item.doClick();
            }
        };

        if (!(item instanceof JMenu)) {
            LibGtk.INSTANCE.g_signal_connect_data(menu.getPointer(), activate, gtkmenu.activate, null, null, 0);
        }

        menukepper.add(gtkmenu);

        return menu;
    }

    GtkWidget createSubmenu(JMenu menu) {
        GtkWidget gmenu = LibGtk.INSTANCE.gtk_menu_new();

        for (int i = 0; i < menu.getMenuComponentCount(); i++) {
            Component e = menu.getMenuComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;

                GtkWidget ss = createSubmenu(sub);
                GtkWidget item = createMenuItem(sub);
                LibGtk.INSTANCE.gtk_menu_item_set_submenu(item, ss);
                LibGtk.INSTANCE.gtk_menu_shell_append(gmenu, item);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                GtkWidget item = createMenuItem(ch);
                LibGtk.INSTANCE.gtk_menu_shell_append(gmenu, item);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                GtkWidget item = createMenuItem(mi);
                LibGtk.INSTANCE.gtk_menu_shell_append(gmenu, item);
            }

            if (e instanceof JPopupMenu.Separator) {
                GtkWidget item = LibGtk.INSTANCE.gtk_separator_menu_item_new();
                LibGtk.INSTANCE.gtk_menu_shell_append(gmenu, item);
            }
        }

        return gmenu;

    }

    void updateMenus() {
        if (gtkmenu != null) {
            gtkmenu.destory();
        }

        menukepper.clear();

        gtkmenu = LibGtk.INSTANCE.gtk_menu_new();

        for (int i = 0; i < menu.getComponentCount(); i++) {
            Component e = menu.getComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;

                GtkWidget ss = createSubmenu(sub);
                GtkWidget item1 = createMenuItem(sub);
                LibGtk.INSTANCE.gtk_menu_item_set_submenu(item1, ss);
                LibGtk.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                GtkWidget item1 = createMenuItem(ch);
                LibGtk.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                GtkWidget item1 = createMenuItem(mi);
                LibGtk.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            }

            if (e instanceof JPopupMenu.Separator) {
                GtkWidget item1 = LibGtk.INSTANCE.gtk_separator_menu_item_new();
                LibGtk.INSTANCE.gtk_menu_shell_append(gtkmenu, item1);
            }
        }
    }

    GtkStatusIcon createGStatusIcon() {
        GtkStatusIcon gicon = LibGtk.INSTANCE.gtk_status_icon_new_from_gicon(convertMenuImage(icon));

        gicon.activate = new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                for (Listener l : Collections.synchronizedCollection(listeners)) {
                    l.mouseLeftClick();
                }
            }
        };

        gicon.popup_menu = new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                LibGtk.INSTANCE.gtk_menu_popup(gtkmenu, null, null, LibGtk.gtk_status_icon_position_menu, data, 1,
                        LibGtk.INSTANCE.gtk_get_current_event_time());
            }
        };

        LibGtk.INSTANCE.g_signal_connect_data(gicon.getPointer(), activate, gicon.activate, null, null, 0);

        LibGtk.INSTANCE.g_signal_connect_data(gicon.getPointer(), popup_menu, gicon.popup_menu, null, null, 0);

        LibGtk.INSTANCE.gtk_status_icon_set_tooltip_text(gicon, title);

        return gicon;
    }

    //
    // public
    //

    public LinuxSysTrayGtk() {
        GtkMessageLoop.inc();
    }

    protected void finalize() throws Throwable {
        super.finalize();

        GtkMessageLoop.dec();

        close();
    }

    @Override
    public void setIcon(Icon icon) {
        this.icon = icon;
    }

    @Override
    public void setTitle(String title) {
        this.title = title;
    }

    GSourceFunc show = new GSourceFunc() {
        @Override
        public boolean invoke(Pointer data) {
            updateMenus();

            if (gtkstatusicon == null)
                gtkstatusicon = createGStatusIcon();

            LibGtk.INSTANCE.gtk_status_icon_set_visible(gtkstatusicon, true);
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

            if (gtkstatusicon != null) {
                LibGtk.INSTANCE.gtk_status_icon_set_from_gicon(gtkstatusicon, convertMenuImage(icon));
                LibGtk.INSTANCE.gtk_status_icon_set_tooltip_text(gtkstatusicon, title);
            }

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
            return false;
        }
    };

    @Override
    public void hide() {
        GtkMessageLoop.invokeLater(hide, null);
    }

    @Override
    public void setMenu(JPopupMenu menu) {
        this.menu = menu;
    }

    @Override
    public void close() {
        if (gtkstatusicon != null) {
            gtkstatusicon.unref();
            gtkstatusicon = null;
        }

        if (gtkmenu != null) {
            gtkmenu.destory();
            gtkmenu = null;
        }
    }

}
