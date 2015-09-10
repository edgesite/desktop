package com.github.axet.desktop.os.mac.cocoa;

import com.github.axet.desktop.os.mac.foundation.Runtime;
import com.sun.jna.Pointer;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSFontDescriptor_Class

public class NSFontDescriptor extends NSObject {

    // AppKit/NSFontDescriptor.m
    public static String NSFontNameAttribute = "NSFontNameAttribute";
    public static String NSFontFamilyAttribute = "NSFontFamilyAttribute";
    public static String NSFontSizeAttribute = "NSFontSizeAttribute";
    public static String NSFontMatrixAttribute = "NSFontMatrixAttribute";
    public static String NSFontCharacterSetAttribute = "NSFontCharacterSetAttribute";
    public static String NSFontTraitsAttribute = "NSFontTraitsAttribute";
    public static String NSFontFaceAttribute = "NSFontFaceAttribute";
    public static String NSFontFixedAdvanceAttribute = "NSFontFixedAdvanceAttribute";
    public static String NSFontVisibleNameAttribute = "NSFontVisibleNameAttribute";

    static Pointer objectForKey = Runtime.INSTANCE.sel_getUid("objectForKey:");

    public NSFontDescriptor(long l) {
        super(l);
    }

    public NSFontDescriptor(Pointer p) {
        this(Pointer.nativeValue(p));
    }

    public NSObject objectForKey(String key) {
        return new NSObject(Runtime.INSTANCE.objc_msgSend(this, objectForKey, new NSString(key)));
    }
}
