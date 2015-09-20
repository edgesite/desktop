package com.github.axet.desktop;

import java.util.HashSet;
import java.util.Set;

import javax.swing.Icon;
import javax.swing.JPopupMenu;

public abstract class DesktopSysTray {

    public interface Listener {
        // Left Click not available in Ubuntu Linux, Mac OS X
        public void mouseLeftClick();
        
        // Double Click not handled in Ubuntu Linux, Mac OS X

        public void mouseLeftDoubleClick();

        // We do not handle right click, because:
        //
        // 1) it is binded to context menu anyway
        //
        // 2) if you call showContextMenu from another java thread, HMENU bugged
        // and you can't use it.
        //
        // 3) Mac OSX does not support showing context menu programmatically
    }

    protected Set<Listener> listeners = new HashSet<Listener>();

    public void addListener(Listener l) {
        listeners.add(l);
    }

    public void removeListener(Listener l) {
        listeners.remove(l);
    }

    public abstract void setIcon(Icon icon);

    /**
     * OSX does not show title on icons.
     * 
     * @param title
     *            set main icon text
     */
    public abstract void setTitle(String title);

    public abstract void show();

    public abstract void update();

    public abstract void hide();

    public abstract void setMenu(JPopupMenu menu);

    public abstract void close();
}
