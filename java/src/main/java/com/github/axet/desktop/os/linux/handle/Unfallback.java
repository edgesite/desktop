package com.github.axet.desktop.os.linux.handle;

import com.sun.jna.Callback;
import com.sun.jna.Pointer;

public interface Unfallback extends Callback {
    void unfallback(Pointer app, GtkStatusIcon icon);
}
