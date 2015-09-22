package com.github.axet.desktop.os.linux.handle;

import com.sun.jna.Callback;
import com.sun.jna.Pointer;

public interface GSourceFunc extends Callback {
    // return false
    boolean invoke(Pointer data);
}
