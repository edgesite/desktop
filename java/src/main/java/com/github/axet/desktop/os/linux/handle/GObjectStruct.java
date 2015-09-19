package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;

public class GObjectStruct extends Structure {
    public class ByValue extends GObjectStruct implements Structure.ByValue {
    }

    public class ByReference extends GObjectStruct implements Structure.ByReference {
    }

    public GTypeInstanceStruct g_type_instance;
    public int ref_count;
    public Pointer qdata;

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("g_type_instance", "ref_count", "qdata");
    }
}
