package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;

public class GTypeInstanceStruct extends Structure {
    public class ByValue extends GTypeInstanceStruct implements Structure.ByValue {
    }

    public class ByReference extends GTypeInstanceStruct implements Structure.ByReference {
    }

    public Pointer g_class;

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("g_class");
    }
}
