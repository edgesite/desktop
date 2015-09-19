package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.NativeLong;
import com.sun.jna.Structure;

public class GTypeClassStruct extends Structure {
    public class ByValue extends GTypeClassStruct implements Structure.ByValue {
    }

    public class ByReference extends GTypeClassStruct implements Structure.ByReference {
    }

    public NativeLong g_type;

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("g_type");
    }
}
