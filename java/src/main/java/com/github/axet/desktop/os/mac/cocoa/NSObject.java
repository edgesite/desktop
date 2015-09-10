package com.github.axet.desktop.os.mac.cocoa;

import com.github.axet.desktop.os.mac.foundation.Runtime;
import com.sun.jna.Pointer;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSObject_Class/index.html#//apple_ref/occ/cl/NSObject

public class NSObject extends Pointer {

    public static Pointer klass = Runtime.INSTANCE.objc_lookUpClass("NSObject");
    static Pointer alloc = Runtime.INSTANCE.sel_getUid("alloc");
    static Pointer retain = Runtime.INSTANCE.sel_getUid("retain");
    static Pointer release = Runtime.INSTANCE.sel_getUid("release");

    public NSObject() {
        super(Runtime.INSTANCE.objc_msgSend(klass, alloc));

        retain();
    }

    public NSObject(Pointer p) {
        super(Pointer.nativeValue(p));

        retain();
    }

    public NSObject(long l) {
        super(l);

        retain();
    }

    protected void finalize() throws Throwable {
        super.finalize();

        release();
    }

    public void retain() {
        Runtime.INSTANCE.objc_msgSend(this, retain);
    }

    public void release() {
        Runtime.INSTANCE.objc_msgSend(this, release);
    }

}