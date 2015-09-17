package com.github.axet.desktop.os.win.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;
import com.sun.jna.platform.win32.WinDef.LONG;

public class LOGFONT extends Structure {

    public static final int LF_FACESIZE = 32;

    public static class ByValue extends LOGFONT implements Structure.ByValue {
    }

    public static class ByReference extends LOGFONT implements Structure.ByReference {
    }

    public LOGFONT() {
    }

    public LOGFONT(Pointer p) {
        super(p);

        read();
    }

    @Override
    protected List<?> getFieldOrder() {
        return Arrays.asList(new String[] { "lfHeight", "lfWidth", "lfEscapement", "lfOrientation", "lfWeight",
                "lfItalic", "lfUnderline", "lfStrikeOut", "lfCharSet", "lfOutPrecision", "lfClipPrecision", "lfQuality",
                "lfPitchAndFamily", "lfFaceName" });
    }

    public LONG lfHeight;
    public LONG lfWidth;
    public LONG lfEscapement;
    public LONG lfOrientation;
    public LONG lfWeight;
    public byte lfItalic;
    public byte lfUnderline;
    public byte lfStrikeOut;
    public byte lfCharSet;
    public byte lfOutPrecision;
    public byte lfClipPrecision;
    public byte lfQuality;
    public byte lfPitchAndFamily;
    public char[] lfFaceName = new char[LF_FACESIZE];
}
