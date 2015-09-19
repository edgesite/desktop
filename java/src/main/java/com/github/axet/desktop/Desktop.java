package com.github.axet.desktop;

import java.io.File;
import java.net.URI;
import java.net.URISyntaxException;

import org.apache.commons.lang.SystemUtils;

import com.github.axet.desktop.os.linux.LinuxFolders;
import com.github.axet.desktop.os.linux.LinuxPower;
import com.github.axet.desktop.os.linux.LinuxSysTrayAppIndicator;
import com.github.axet.desktop.os.linux.LinuxSysTrayGtk;
import com.github.axet.desktop.os.mac.OSXFolders;
import com.github.axet.desktop.os.mac.OSXPower;
import com.github.axet.desktop.os.mac.OSXSysTray;
import com.github.axet.desktop.os.win.WindowsFolders;
import com.github.axet.desktop.os.win.WindowsPowerVista;
import com.github.axet.desktop.os.win.WindowsPowerXP;
import com.github.axet.desktop.os.win.WindowsSysTray;
import com.sun.jna.Platform;

public abstract class Desktop {

    static DesktopFolders desktopFolders = null;
    static DesktopSysTray desktopSysTray = null;
    static DesktopPower desktopPower = null;

    //
    // Desktop Folders
    //

    public static File getHomeFolder() {
        return getDesktopFolders().getHome();
    }

    public static File getDocumentsFolder() {
        return getDesktopFolders().getDocuments();
    }

    public static File getAppDataFolder() {
        return getDesktopFolders().getAppData();
    }

    public static File getDesktopFolder() {
        return getDesktopFolders().getDesktop();
    }

    public static File getDownloadsFolder() {
        return getDesktopFolders().getDownloads();
    }

    public static DesktopFolders getDesktopFolders() {
        if (desktopFolders == null) {
            if (com.sun.jna.Platform.isWindows())
                desktopFolders = new WindowsFolders();

            if (com.sun.jna.Platform.isMac())
                desktopFolders = new OSXFolders();

            if (com.sun.jna.Platform.isLinux())
                desktopFolders = new LinuxFolders();

            if (desktopFolders == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopFolders;
    }

    //
    // Browser
    //

    public static void browserOpenURI(String s) {
        try {
            URI uri;
            uri = new URI(s);
            java.awt.Desktop desktop = java.awt.Desktop.isDesktopSupported() ? java.awt.Desktop.getDesktop() : null;
            if (desktop != null && desktop.isSupported(java.awt.Desktop.Action.BROWSE)) {
                try {
                    desktop.browse(uri);
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        } catch (URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    //
    // SysTray
    //

    public static DesktopSysTray getDesktopSysTray() {
        if (desktopSysTray == null) {
            if (com.sun.jna.Platform.isWindows()) {
                desktopSysTray = new WindowsSysTray();
            }

            if (com.sun.jna.Platform.isMac())
                desktopSysTray = new OSXSysTray();

            if (Platform.isLinux())
                desktopSysTray = new LinuxSysTrayAppIndicator();

            if (desktopSysTray == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopSysTray;
    }

    //
    // Power
    //

    public static DesktopPower getDesktopPower() {
        if (desktopPower == null) {
            if (SystemUtils.IS_OS_WINDOWS) {
                if (SystemUtils.IS_OS_WINDOWS_XP)
                    desktopPower = new WindowsPowerXP();
                else
                    desktopPower = new WindowsPowerVista();
            }

            if (com.sun.jna.Platform.isMac())
                desktopPower = new OSXPower();

            if (Platform.isLinux())
                desktopPower = new LinuxPower();

            if (desktopPower == null)
                throw new RuntimeException("OS not supported");
        }

        return desktopPower;
    }

}
