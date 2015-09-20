package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;

public class GtkWidget extends GObject {
    public GtkWidget() {
    }

    public GtkWidget(Pointer p) {
        super(p);
    }

    public void destory() {
        LibGtk.INSTANCE.gtk_widget_destroy(getPointer());
    }
}
