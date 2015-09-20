package com.github.axet.desktop.os.linux.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;

public class AppIndicatorClassStruct extends Structure {
    public class ByReference extends AppIndicatorClassStruct implements Structure.ByReference {
    }

    public GObjectClassStruct parent_class;

    public Pointer new_icon;
    public Pointer new_attention_icon;
    public Pointer new_status;
    public Pointer new_icon_theme;
    public Pointer new_label;
    public Pointer connection_changed;
    public Pointer scroll_event;
    public Pointer app_indicator_reserved_ats;
    public Fallback fallback;
    public Unfallback unfallback;
    public Pointer app_indicator_reserved_1;
    public Pointer app_indicator_reserved_2;
    public Pointer app_indicator_reserved_3;
    public Pointer app_indicator_reserved_4;
    public Pointer app_indicator_reserved_5;
    public Pointer app_indicator_reserved_6;

    public AppIndicatorClassStruct() {
    }

    public AppIndicatorClassStruct(Pointer p) {
        super(p);
        useMemory(p);
        read();
    }

    @Override
    protected List<String> getFieldOrder() {
        return Arrays.asList("parent_class", "new_icon", "new_attention_icon", "new_status", "new_icon_theme",
                "new_label", "connection_changed", "scroll_event", "app_indicator_reserved_ats", "fallback",
                "unfallback", "app_indicator_reserved_1", "app_indicator_reserved_2", "app_indicator_reserved_3",
                "app_indicator_reserved_4", "app_indicator_reserved_5", "app_indicator_reserved_6");
    }
}
