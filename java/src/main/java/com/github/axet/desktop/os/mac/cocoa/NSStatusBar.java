package com.github.axet.desktop.os.mac.cocoa;

import com.github.axet.desktop.os.mac.foundation.Runtime;
import com.sun.jna.Pointer;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSStatusBar_Class

public class NSStatusBar extends NSObject {

    public final static int NSVariableStatusItemLength = -1;
    public final static int NSSquareStatusItemLength = -2;

    static Pointer klass = Runtime.INSTANCE.objc_lookUpClass("NSStatusBar");
    static Pointer systemStatusBar = Runtime.INSTANCE.sel_getUid("systemStatusBar");
    static Pointer statusItemWithLength = Runtime.INSTANCE.sel_getUid("statusItemWithLength:");
    static Pointer removeStatusItem = Runtime.INSTANCE.sel_getUid("removeStatusItem:");
    static Pointer thickness = Runtime.INSTANCE.sel_getUid("thickness");

    public static NSStatusBar systemStatusBar() {
        return new NSStatusBar(Runtime.INSTANCE.objc_msgSend(klass, systemStatusBar));
    }

    public NSStatusBar(long l) {
        super(l);
    }

    public NSStatusBar(Pointer p) {
        this(Pointer.nativeValue(p));
    }

    public NSStatusItem statusItemWithLength(double i) {
        return new NSStatusItem(Runtime.INSTANCE.objc_msgSend(this, statusItemWithLength, i));
    }

    public void removeStatusItem(NSStatusItem i) {
        Runtime.INSTANCE.objc_msgSend(this, removeStatusItem, i);
    }

    public CGFloat thickness() {
        return new CGFloat(Runtime.INSTANCE.objc_msgSend(this, thickness));
    }
}