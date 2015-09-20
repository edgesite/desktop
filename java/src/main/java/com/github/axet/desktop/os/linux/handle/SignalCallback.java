package com.github.axet.desktop.os.linux.handle;

import com.sun.jna.Callback;
import com.sun.jna.Pointer;

public interface SignalCallback extends Callback {
    void signal(Pointer data);
}
