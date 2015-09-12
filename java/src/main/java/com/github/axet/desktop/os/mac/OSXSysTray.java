package com.github.axet.desktop.os.mac;

import java.awt.AlphaComposite;
import java.awt.Component;
import java.awt.Graphics2D;
import java.awt.image.BufferedImage;
import java.util.ArrayList;

import javax.swing.Icon;
import javax.swing.JCheckBoxMenuItem;
import javax.swing.JMenu;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.Utils;
import com.github.axet.desktop.os.mac.cocoa.NSCell;
import com.github.axet.desktop.os.mac.cocoa.NSFont;
import com.github.axet.desktop.os.mac.cocoa.NSFontDescriptor;
import com.github.axet.desktop.os.mac.cocoa.NSImage;
import com.github.axet.desktop.os.mac.cocoa.NSMenu;
import com.github.axet.desktop.os.mac.cocoa.NSMenuItem;
import com.github.axet.desktop.os.mac.cocoa.NSNumber;
import com.github.axet.desktop.os.mac.cocoa.NSObject;
import com.github.axet.desktop.os.mac.cocoa.NSStatusBar;
import com.github.axet.desktop.os.mac.cocoa.NSStatusItem;
import com.github.axet.desktop.os.mac.cocoa.NSString;

public class OSXSysTray extends DesktopSysTray {
    // keep title to update it when icon reshown
    NSImage icon;
    // keep title to update it when icon reshown
    String title;
    // kepp to abble to rebuild menu
    JPopupMenu menu;
    // keep reference to be able to remove statusbar icon
    NSStatusItem statusItem;
    // prevent action to be gc(), since it has two pointers to NS object and
    // Java object. nsobject keept by Apple and here is no one who keeps java
    // reference.
    ArrayList<NSObject> actionKeeper = new ArrayList<NSObject>();
    NSStatusBar statusbar;

    public OSXSysTray() {
        // init menubar font, to get proper font sizes
        statusbar = NSStatusBar.systemStatusBar();
    }

    @Override
    public void setIcon(Icon icon) {
        this.icon = convertTrayIcon(icon);
    }

    static NSImage convertTrayIcon(Icon i) {
        BufferedImage icon = Utils.createBitmap(i);

        NSFont f = NSFont.menuBarFontOfSize(0);
        int menubarHeigh = new NSNumber(f.fontDescriptor().objectForKey(NSFontDescriptor.NSFontSizeAttribute))
                .intValue();

        BufferedImage scaledImage = new BufferedImage(menubarHeigh, menubarHeigh, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = scaledImage.createGraphics();
        g.setComposite(AlphaComposite.getInstance(AlphaComposite.SRC_OVER));
        g.drawImage(icon, 0, 0, menubarHeigh, menubarHeigh, null);
        g.dispose();

        return new NSImage(scaledImage);
    }

    static NSImage convertMenuIcon(Icon icon) {
        BufferedImage img = Utils.createBitmap(icon);

        NSFont f = NSFont.menuFontOfSize(0);
        int menubarHeigh = new NSNumber(f.fontDescriptor().objectForKey(NSFontDescriptor.NSFontSizeAttribute))
                .intValue();

        BufferedImage scaledImage = new BufferedImage(menubarHeigh, menubarHeigh, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = scaledImage.createGraphics();
        g.setComposite(AlphaComposite.getInstance(AlphaComposite.SRC_OVER));
        g.drawImage(img, 0, 0, menubarHeigh, menubarHeigh, null);
        g.dispose();

        return new NSImage(scaledImage);
    }

    @Override
    public void setTitle(String title) {
        this.title = title;
    }

    @Override
    public void show() {
        updateMenus();
    }

    void updateMenus() {
        if (statusItem == null) {
            statusItem = statusbar.statusItemWithLength(NSStatusBar.NSVariableStatusItemLength);
        }

        statusItem.setToolTip(title);

        actionKeeper.clear();

        NSMenu m = new NSMenu();

        for (int i = 0; i < menu.getComponentCount(); i++) {
            Component e = menu.getComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;
                NSMenu hsub = createSubmenu(sub);

                NSImage bm = null;
                if (sub.getIcon() != null)
                    bm = convertMenuIcon(sub.getIcon());

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(sub.getText()));
                item.setImage(bm);
                item.setSubmenu(hsub);
                m.addItem(item);
            } else if (e instanceof JCheckBoxMenuItem) {
                JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                NSImage bm = null;
                if (ch.getIcon() != null)
                    bm = convertMenuIcon(ch.getIcon());

                OSXSysTrayAction action = new OSXSysTrayAction(ch);
                actionKeeper.add(action);

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(ch.getText()));
                item.setImage(bm);
                item.setEnabled(ch.isEnabled());
                item.setState(ch.getState() ? NSCell.NSCellStateValue.NSOnState : NSCell.NSCellStateValue.NSOffState);
                item.setTarget(action);
                item.setAction(OSXSysTrayAction.action);
                m.addItem(item);
            } else if (e instanceof JMenuItem) {
                JMenuItem mi = (JMenuItem) e;

                NSImage bm = null;
                if (mi.getIcon() != null)
                    bm = convertMenuIcon(mi.getIcon());

                OSXSysTrayAction action = new OSXSysTrayAction(mi);
                actionKeeper.add(action);

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(mi.getText()));
                item.setImage(bm);
                item.setEnabled(mi.isEnabled());
                item.setTarget(action);
                item.setAction(OSXSysTrayAction.action);
                m.addItem(item);
            }

            if (e instanceof JPopupMenu.Separator) {
                m.addItem(NSMenuItem.separatorItem());
            }
        }

        m.setAutoenablesItems(false);
        statusItem.setImage(icon);
        statusItem.setHighlightMode(true);
        statusItem.setMenu(m);
    }

    NSMenu createSubmenu(JMenu menu) {
        NSMenu m = new NSMenu();

        for (int i = 0; i < menu.getMenuComponentCount(); i++) {
            Component e = menu.getMenuComponent(i);

            if (e instanceof JMenu) {
                JMenu sub = (JMenu) e;
                NSMenu hsub2 = createSubmenu(sub);

                NSImage bm = null;
                if (sub.getIcon() != null)
                    bm = convertMenuIcon(sub.getIcon());

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(sub.getText()));
                item.setImage(bm);
                item.setSubmenu(hsub2);
                m.addItem(item);
            } else if (e instanceof JCheckBoxMenuItem) {
                JCheckBoxMenuItem ch = (JCheckBoxMenuItem) e;

                NSImage bm = null;
                if (ch.getIcon() != null)
                    bm = convertMenuIcon(ch.getIcon());

                OSXSysTrayAction action = new OSXSysTrayAction(ch);
                actionKeeper.add(action);

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(ch.getText()));
                item.setImage(bm);
                item.setEnabled(ch.isEnabled());
                item.setState(ch.getState() ? 1 : 0);
                item.setTarget(action);
                item.setAction(OSXSysTrayAction.action);
                m.addItem(item);
            } else if (e instanceof JMenuItem) {
                JMenuItem mi = (JMenuItem) e;

                NSImage bm = null;
                if (mi.getIcon() != null)
                    bm = convertMenuIcon(mi.getIcon());

                OSXSysTrayAction action = new OSXSysTrayAction(mi);
                actionKeeper.add(action);

                NSMenuItem item = new NSMenuItem();
                item.setTitle(new NSString(mi.getText()));
                item.setImage(bm);
                item.setEnabled(mi.isEnabled());
                item.setTarget(action);
                item.setAction(OSXSysTrayAction.action);
                m.addItem(item);

            }

            if (e instanceof JPopupMenu.Separator) {
                m.addItem(NSMenuItem.separatorItem());
            }
        }

        return m;
    }

    @Override
    public void update() {
        updateMenus();
    }

    @Override
    public void hide() {
        if (statusItem != null) {
            statusbar.removeStatusItem(statusItem);
            statusItem = null;
            actionKeeper.clear();
        }
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
