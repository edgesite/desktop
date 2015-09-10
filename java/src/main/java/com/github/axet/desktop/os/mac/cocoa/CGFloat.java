package com.github.axet.desktop.os.mac.cocoa;

public class CGFloat {
    double d;

    public CGFloat(long l) {
        d = Double.longBitsToDouble(l);
    }

    public double Double() {
        return d;
    }

    public int Int() {
        return (int) d;
    }
}
