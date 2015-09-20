package com.github.axet.desktop.os.linux;

import javax.swing.Icon;
import javax.swing.JPopupMenu;

import com.github.axet.desktop.DesktopSysTray;

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

public class LinuxSysTrayXEvent extends DesktopSysTray {

    public LinuxSysTrayXEvent() {
    }

    protected void finalize() throws Throwable {
        super.finalize();
    }

    @Override
    public void setIcon(Icon icon) {
    }

    @Override
    public void setTitle(String title) {
    }

    @Override
    public void show() {
    }

    @Override
    public void update() {
    }

    @Override
    public void hide() {
    }

    @Override
    public void setMenu(JPopupMenu menu) {
    }

    @Override
    public void close() {
        hide();
    }

}
