package com.github.axet.desktop;

import java.awt.AlphaComposite;
import java.awt.Graphics2D;
import java.awt.image.BufferedImage;
import java.io.IOException;
import java.io.InputStream;

import javax.imageio.ImageIO;
import javax.swing.Icon;
import javax.swing.ImageIcon;

public class Utils {

    public static BufferedImage createBitmap(Icon icon) {
        BufferedImage bi = new BufferedImage(icon.getIconWidth(), icon.getIconHeight(), BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = bi.createGraphics();
        g.setComposite(AlphaComposite.getInstance(AlphaComposite.SRC_OVER));
        icon.paintIcon(null, g, 0, 0);
        g.dispose();
        return bi;
    }

    public static Icon loadIcon(String path) {
        try {
            InputStream is = Utils.class.getResourceAsStream(path);
            if (is == null) {
                throw new RuntimeException("resource file not found " + path);
            }
            BufferedImage bi = ImageIO.read(is);
            if (bi == null) {
                throw new RuntimeException("wrong format " + path);
            }
            return new ImageIcon(bi);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

}
