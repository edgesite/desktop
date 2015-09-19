package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibAppIndicator;
import com.sun.jna.Pointer;
import com.sun.jna.PointerType;

public class GObject extends PointerType {
    public GObject() {
    }

    public GObject(Pointer p) {
        super(p);
    }

    protected void finalize() throws Throwable {
        super.finalize();
    }

    public void ref() {
        LibAppIndicator.INSTANCE.g_object_ref(this);
    }

    public void unref() {
        LibAppIndicator.INSTANCE.g_object_unref(this);
    }
}
