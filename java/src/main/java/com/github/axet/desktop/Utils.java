package com.github.axet.desktop;

import java.awt.AlphaComposite;
import java.awt.Graphics2D;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.io.IOException;

import javax.imageio.ImageIO;
import javax.swing.Icon;

public class Utils {

    public static BufferedImage createBitmap(Icon icon) {
        BufferedImage bi = new BufferedImage(icon.getIconWidth(), icon.getIconHeight(), BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = bi.createGraphics();
        g.setComposite(AlphaComposite.getInstance(AlphaComposite.SRC_OVER));
        icon.paintIcon(null, g, 0, 0);
        g.dispose();
        return bi;
    }

    public static byte[] BufferedImage2Bytes(BufferedImage b) {
        try {
            ByteArrayOutputStream os = new ByteArrayOutputStream();
            ImageIO.write(b, "png", os);
            os.flush();
            return os.toByteArray();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
