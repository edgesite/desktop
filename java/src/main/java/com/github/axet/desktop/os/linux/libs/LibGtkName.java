package com.github.axet.desktop.os.linux.libs;

import com.github.axet.desktop.DesktopSysTray;
import com.github.axet.desktop.os.linux.LinuxSysTrayAppIndicator;
import com.github.axet.desktop.os.linux.LinuxSysTrayGtk;
import com.sun.jna.NativeLibrary;

public class LibGtkName {

    static String NAME = null;

    static boolean APPINDICATOR = true;
    static boolean GTK = true;

    public static DesktopSysTray createSysTray() {

        if (APPINDICATOR) {
            String ss[] = new String[] { "appindicator3", "appindicator" };

            for (String s : ss) {
                try {
                    NativeLibrary.getInstance(s);
                    NAME = s;
                    return new LinuxSysTrayAppIndicator();
                } catch (java.lang.UnsatisfiedLinkError e) {
                    ;
                }
            }
        }

        if (GTK) {
            NAME = "Gtk3";
            return new LinuxSysTrayGtk();
        }

        throw new RuntimeException("library not found");
    }

    public static String getName() {
        return NAME;
    }

}
