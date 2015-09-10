package com.github.axet.desktop.os.mac.cocoa;

import com.github.axet.desktop.os.mac.foundation.Runtime;
import com.sun.jna.Pointer;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSNumber_Class/

public class NSNumber extends NSObject {
    static Pointer floatValue = Runtime.INSTANCE.sel_getUid("floatValue");
    static Pointer intValue = Runtime.INSTANCE.sel_getUid("intValue");
    static Pointer doubleValue = Runtime.INSTANCE.sel_getUid("doubleValue");
    static Pointer stringValue = Runtime.INSTANCE.sel_getUid("stringValue");

    public NSNumber(long l) {
        super(l);
    }

    public NSNumber(Pointer p) {
        this(Pointer.nativeValue(p));
    }

    public float floatValue() {
        long f = Runtime.INSTANCE.objc_msgSend(this, floatValue);
        return Float.floatToIntBits(f);
    }

    public int intValue() {
        return (int) Runtime.INSTANCE.objc_msgSend(this, intValue);
    }

    public double doubleValue() {
        return Double.longBitsToDouble(Runtime.INSTANCE.objc_msgSend(this, doubleValue));
    }

    public String toString() {
        NSString s = new NSString(Runtime.INSTANCE.objc_msgSend(this, stringValue));
        return s.toString();
    }
}
