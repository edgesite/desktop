package com.github.axet.desktop.os.win.libs;

import com.github.axet.desktop.os.win.handle.COLORREF;
import com.github.axet.desktop.os.win.handle.LOGFONT;
import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.platform.win32.WinDef.HDC;
import com.sun.jna.platform.win32.WinDef.HFONT;
import com.sun.jna.platform.win32.WinDef.RECT;
import com.sun.jna.platform.win32.WinUser.SIZE;
import com.sun.jna.win32.W32APIOptions;

public interface GDI32Ex extends Library {

    public static final int ETO_OPAQUE = 2;
    public static final int SRCCOPY = 0xCC0020;

    static GDI32Ex INSTANCE = (GDI32Ex) Native.loadLibrary("GDI32", GDI32Ex.class, W32APIOptions.DEFAULT_OPTIONS);

    // http://msdn.microsoft.com/en-us/library/windows/desktop/dd144938(v=vs.85).aspx
    boolean GetTextExtentPoint32(HDC hdc, String lpString, int c, SIZE lpSize);

    // http://msdn.microsoft.com/en-us/library/windows/desktop/dd145093(v=vs.85).aspx
    COLORREF SetTextColor(HDC hdc, COLORREF crColor);

    // http://msdn.microsoft.com/en-us/library/windows/desktop/dd162964(v=vs.85).aspx
    COLORREF SetBkColor(HDC hdc, COLORREF crColor);

    // http://msdn.microsoft.com/en-us/library/windows/desktop/dd162713(v=vs.85).aspx
    boolean ExtTextOut(HDC hdc, int X, int Y, int fuOptions, RECT lprc, String lpString, int cbCount, int[] lpDx);

    // http://msdn.microsoft.com/en-us/library/windows/desktop/dd183370(v=vs.85).aspx
    boolean BitBlt(HDC hdcDest, int nXDest, int nYDest, int nWidth, int nHeight, HDC hdcSrc, int nXSrc, int nYSrc,
            int dwRop);

    // https://msdn.microsoft.com/ru-ru/library/windows/desktop/dd183500(v=vs.85).aspx
    HFONT CreateFontIndirect(LOGFONT l);

}
