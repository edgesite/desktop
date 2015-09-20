package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;
import com.sun.jna.PointerType;

public class GBytes extends PointerType {
    public GBytes() {
    }

    public GBytes(Pointer p) {
        super(p);
    }

    protected void finalize() throws Throwable {
        System.out.println(getClass().getSimpleName());
        
        unref();

        super.finalize();
    }

    public void unref() {
        LibGtk.INSTANCE.g_bytes_unref(this);
    }
}
