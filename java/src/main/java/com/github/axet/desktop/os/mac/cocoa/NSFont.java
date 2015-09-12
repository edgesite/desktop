package com.github.axet.desktop.os.mac.cocoa;

import com.github.axet.desktop.os.mac.foundation.Runtime;
import com.sun.jna.Pointer;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSFont_Class

public class NSFont extends NSObject {

    static Pointer klass = Runtime.INSTANCE.objc_lookUpClass("NSFont");
    static Pointer menuFontOfSize = Runtime.INSTANCE.sel_getUid("menuFontOfSize:");
    static Pointer menuBarFontOfSize = Runtime.INSTANCE.sel_getUid("menuBarFontOfSize:");
    static Pointer pointSize = Runtime.INSTANCE.sel_getUid("pointSize");
    static Pointer fontName = Runtime.INSTANCE.sel_getUid("fontName");
    static Pointer displayName = Runtime.INSTANCE.sel_getUid("displayName");
    static Pointer fontDescriptor = Runtime.INSTANCE.sel_getUid("fontDescriptor");

    public static NSFont menuFontOfSize(double size) {
        return new NSFont(Runtime.INSTANCE.objc_msgSend(klass, menuFontOfSize, size));
    }

    public static NSFont menuBarFontOfSize(double size) {
        return new NSFont(Runtime.INSTANCE.objc_msgSend(klass, menuBarFontOfSize, size));
    }

    public NSFont(long l) {
        super(l);
    }

    public NSFont(Pointer p) {
        this(Pointer.nativeValue(p));
    }

    public CGFloat pointSize() {
        return new CGFloat(Runtime.INSTANCE.objc_msgSend(this, pointSize));
    }

    public NSString fontName() {
        return new NSString(Runtime.INSTANCE.objc_msgSend(this, fontName));
    }

    public NSString displayName() {
        return new NSString(Runtime.INSTANCE.objc_msgSend(this, displayName));
    }

    public NSFontDescriptor fontDescriptor() {
        return new NSFontDescriptor(Runtime.INSTANCE.objc_msgSend(this, fontDescriptor));
    }
}
