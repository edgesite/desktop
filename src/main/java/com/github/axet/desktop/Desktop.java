package com.github.axet.desktop;

import org.apache.commons.lang.SystemUtils;

import com.github.axet.desktop.os.LinuxFolders;
import com.github.axet.desktop.os.mac.OSXFolders;
import com.github.axet.desktop.os.mac.OSXPower;
import com.github.axet.desktop.os.mac.OSXSysTray;
import com.github.axet.desktop.os.win.WindowsFolders;
import com.github.axet.desktop.os.win.WindowsPowerVista;
import com.github.axet.desktop.os.win.WindowsPowerXP;
import com.github.axet.desktop.os.win.WindowsSysTray;

public abstract class Desktop {

    static DesktopFolders desktopFolders = null;
    static DesktopSysTray desktopSysTray = null;
    static DesktopPower desktopPower = null;

    public static DesktopFolders getDesktopFolders() {
        if (desktopFolders == null) {
            if (com.sun.jna.Platform.isWindows()) {
                desktopFolders = new WindowsFolders();
            }

            if (com.sun.jna.Platform.isMac()) {
                desktopFolders = new OSXFolders();
            }

            if (com.sun.jna.Platform.isLinux()) {
                desktopFolders = new LinuxFolders();
            }

            if (desktopFolders == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopFolders;
    }

    public static DesktopSysTray getDesktopSysTray() {
        if (desktopSysTray == null) {
            if (com.sun.jna.Platform.isWindows()) {
                desktopSysTray = new WindowsSysTray();
            }

            if (com.sun.jna.Platform.isMac()) {
                desktopSysTray = new OSXSysTray();
            }

            if (desktopSysTray == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopSysTray;
    }

    public static DesktopPower getDesktopPower() {
        if (desktopPower == null) {
            if (SystemUtils.IS_OS_WINDOWS) {
                if (SystemUtils.IS_OS_WINDOWS_XP)
                    desktopPower = new WindowsPowerXP();
                else
                    desktopPower = new WindowsPowerVista();
            }

            if (com.sun.jna.Platform.isMac()) {
                desktopPower = new OSXPower();
            }

            if (desktopPower == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopPower;
    }

}
