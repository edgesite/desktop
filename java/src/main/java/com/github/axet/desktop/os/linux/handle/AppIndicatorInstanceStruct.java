package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;

public class AppIndicatorInstanceStruct extends Structure {
    public GObjectStruct parent;
    public Pointer priv;

    public AppIndicatorInstanceStruct() {

    }

    public AppIndicatorInstanceStruct(Pointer p) {
        super(p);
        read();
    }

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("parent", "priv");
    }
}
