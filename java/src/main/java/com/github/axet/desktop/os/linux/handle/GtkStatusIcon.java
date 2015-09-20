package com.github.axet.desktop.os.linux.handle;

import com.sun.jna.Pointer;

public class GtkStatusIcon extends GObject {
    
    public SignalCallback activate;
    public SignalCallback popup_menu;
    
    public GtkStatusIcon() {

    }

    public GtkStatusIcon(Pointer p) {
        super(p);
    }
}
