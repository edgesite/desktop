package com.github.axet.desktop.os.linux.libs;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.os.linux.LinuxSysTrayAppIndicator;
import com.github.axet.desktop.os.linux.LinuxSysTrayGtk;
import com.sun.jna.NativeLibrary;

public class LibGtkName {

    // set those variables to control what linux systray algorithm will be used

    public static boolean APPINDICATOR = true;
    public static boolean GTK = true;

    static String NAME = null;

    public static DesktopSysTray createSysTray() {

        if (APPINDICATOR) {
            String ss[] = new String[] { "appindicator3", "appindicator" };

            for (String s : ss) {
                try {
                    NativeLibrary.getInstance(s);
                    NAME = s;
                    return new LinuxSysTrayAppIndicator();
                } catch (java.lang.UnsatisfiedLinkError e) {
                }
            }
        }

        if (GTK) {
            String ss[] = new String[] { "gtk-3", "gdk-x11-2.0" };

            for (String s : ss) {
                try {
                    NativeLibrary.getInstance(s);
                    NAME = s;
                    return new LinuxSysTrayGtk();
                } catch (java.lang.UnsatisfiedLinkError e) {
                }
            }
        }

        throw new RuntimeException("no compatible trayicon library found");
    }

    public static String getName() {
        return NAME;
    }

}
