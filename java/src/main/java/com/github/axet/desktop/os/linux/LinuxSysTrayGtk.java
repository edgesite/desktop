package com.github.axet.desktop.os.linux;

import java.awt.CheckboxMenuItem;
import java.awt.Component;
import java.awt.Menu;
import java.awt.MenuItem;
import java.awt.PopupMenu;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.ItemEvent;
import java.awt.event.ItemListener;
import java.awt.image.BufferedImage;

import javax.swing.Icon;
import javax.swing.JCheckBoxMenuItem;
import javax.swing.JMenu;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.Utils;
import com.github.axet.desktop.os.linux.handle.GBytes;
import com.github.axet.desktop.os.linux.handle.GIcon;
import com.github.axet.desktop.os.linux.handle.GtkStatusIcon;
import com.github.axet.desktop.os.linux.handle.GtkWidget;
import com.github.axet.desktop.os.linux.handle.SignalCallback;
import com.github.axet.desktop.os.linux.libs.LibAppIndicator;
import com.github.axet.desktop.os.linux.libs.Gtk3;
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

public class LinuxSysTrayGtk extends DesktopSysTray {
    PopupMenu popup;
    JPopupMenu menu;
    String title;

    BufferedImage icon;

    public LinuxSysTrayGtk() {
    }

    @Override
    public void setIcon(Icon icon) {
        this.icon = Utils.createBitmap(icon);
    }

    @Override
    public void setTitle(String title) {
        this.title = title;
    }

    @Override
    public void show() {
        Gtk3.INSTANCE.gtk_init(0, null);

        final GtkWidget menu = Gtk3.INSTANCE.gtk_menu_new();
        GtkWidget item = Gtk3.INSTANCE.gtk_menu_item_new_with_label("Item1");
        Gtk3.INSTANCE.g_signal_connect_data(item.getPointer(), "activate", new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                System.out.println("item");
            }
        }, null, null, 0);
        Gtk3.INSTANCE.gtk_menu_shell_append(menu, item);
        Gtk3.INSTANCE.gtk_widget_show(item);

        byte[] buf = Utils.BufferedImage2Bytes(icon);

        GBytes bg = Gtk3.INSTANCE.g_bytes_new(buf, buf.length);

        GIcon g = Gtk3.INSTANCE.g_bytes_icon_new(bg);
        GtkStatusIcon icon = Gtk3.INSTANCE.gtk_status_icon_new_from_gicon(g);

        Gtk3.INSTANCE.gtk_status_icon_set_visible(icon, true);

        Gtk3.INSTANCE.g_signal_connect_data(icon.getPointer(), "activate", new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                System.out.println("lclick");
            }
        }, null, null, 0);

        Gtk3.INSTANCE.g_signal_connect_data(icon.getPointer(), "popup-menu", new SignalCallback() {
            @Override
            public void signal(Pointer data) {
                System.out.println("rclick");
                Function gtk_status_icon_position_menu = Function.getFunction("appindicator",
                        "gtk_status_icon_position_menu");
                int time = Gtk3.INSTANCE.gtk_get_current_event_time();
                Gtk3.INSTANCE.gtk_menu_popup(menu, null, null, gtk_status_icon_position_menu, data, 1, time);
            }
        }, null, null, 0);

        new Thread(new Runnable() {
            
            @Override
            public void run() {
                Pointer mainloop = Gtk3.INSTANCE.g_main_loop_new(null, false);
                Gtk3.INSTANCE.g_main_loop_run(mainloop);
            }
        }).start();
    }

    @Override
    public void update() {
        updateMenus();
    }

    void updateMenus() {
        popup = new PopupMenu();

        for (int i = 0; i < menu.getComponentCount(); i++) {
            Component e = menu.getComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;
                Menu ss = createSubmenu(sub);
                popup.add(ss);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                final CheckboxMenuItem mm = new CheckboxMenuItem(ch.getText(), ch.getState());
                mm.addItemListener(new ItemListener() {
                    @Override
                    public void itemStateChanged(ItemEvent e) {
                        ch.doClick();
                        updateMenus();
                    }
                });
                popup.add(mm);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                final MenuItem mm = new MenuItem(mi.getText());
                mm.addActionListener(new ActionListener() {
                    @Override
                    public void actionPerformed(ActionEvent e) {
                        mi.doClick();
                    }
                });
                popup.add(mm);
            }

            if (e instanceof JPopupMenu.Separator) {
                popup.insertSeparator(popup.getItemCount());
            }
        }
    }

    Menu createSubmenu(JMenu menu) {
        Menu popup = new Menu(menu.getText());

        for (int i = 0; i < menu.getMenuComponentCount(); i++) {
            Component e = menu.getMenuComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;
                Menu ss = createSubmenu(sub);
                popup.add(ss);
            } else if (e instanceof JCheckBoxMenuItem) {
                final JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                final CheckboxMenuItem mm = new CheckboxMenuItem(ch.getText());
                mm.addItemListener(new ItemListener() {
                    @Override
                    public void itemStateChanged(ItemEvent e) {
                        ch.doClick();
                        updateMenus();
                    }
                });
                popup.add(mm);
            } else if (e instanceof JMenuItem) {
                final JMenuItem mi = (JMenuItem) e;

                final MenuItem mm = new MenuItem(mi.getText());
                mm.addActionListener(new ActionListener() {
                    @Override
                    public void actionPerformed(ActionEvent e) {
                        mi.doClick();
                    }
                });
                popup.add(mm);
            }

            if (e instanceof JPopupMenu.Separator) {
                popup.insertSeparator(popup.getItemCount());
            }
        }

        return popup;

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
