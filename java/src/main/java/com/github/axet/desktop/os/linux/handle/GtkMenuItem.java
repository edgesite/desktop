package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;

public class GtkMenuItem extends GObject {
    public SignalCallback activate;
    
    public GtkMenuItem() {
    }

    public GtkMenuItem(Pointer p) {
        super(p);
    }

    public void destory() {
        LibGtk.INSTANCE.gtk_widget_destroy(getPointer());
    }
}
