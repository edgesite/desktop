package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;
import com.sun.jna.PointerType;

public class GObject extends PointerType {
    public GObject() {
    }

    public GObject(Pointer p) {
        super(p);
        
        ref();
    }

    protected void finalize() throws Throwable {
        super.finalize();
        
        unref();
    }

    public void ref() {
        LibGtk.INSTANCE.g_object_ref(getPointer());
    }

    public void unref() {
        LibGtk.INSTANCE.g_object_unref(getPointer());
    }
}
