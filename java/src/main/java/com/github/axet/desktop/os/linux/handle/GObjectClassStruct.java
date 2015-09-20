package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.NativeLong;
import com.sun.jna.Pointer;
import com.sun.jna.Structure;

public class GObjectClassStruct extends Structure {
    public class ByValue extends GObjectClassStruct implements Structure.ByValue {
    }

    public class ByReference extends GObjectClassStruct implements Structure.ByReference {
    }

    public GTypeClassStruct g_type_class;
    public Pointer construct_properties;
    public Pointer constructor;
    public Pointer set_property;
    public Pointer get_property;
    public Pointer dispose;
    public Pointer finalize;
    public Pointer dispatch_properties_changed;
    public Pointer notify;
    public Pointer constructed;
    public NativeLong flags;
    public Pointer dummy1;
    public Pointer dummy2;
    public Pointer dummy3;
    public Pointer dummy4;
    public Pointer dummy5;
    public Pointer dummy6;

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("g_type_class", "construct_properties", "constructor", "set_property", "get_property",
                "dispose", "finalize", "dispatch_properties_changed", "notify", "constructed", "flags", "dummy1",
                "dummy2", "dummy3", "dummy4", "dummy5", "dummy6");
    }
}