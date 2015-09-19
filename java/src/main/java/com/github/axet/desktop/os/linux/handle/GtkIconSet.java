package com.github.axet.desktop.os.linux.handle;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.HashMap;

import javax.imageio.ImageIO;
import javax.swing.Icon;

import org.apache.commons.io.FileUtils;

import com.github.axet.desktop.Utils;

public class GtkIconSet {
    File tmp;
    HashMap<Icon, String> map = new HashMap<Icon, String>();

    public GtkIconSet() {
        try {
            tmp = Files.createTempDirectory("systrayicon").toFile();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public String getPath() {
        return tmp.getAbsolutePath();
    }

    public String getIcon(Icon image) {
        return map.get(image);
    }

    public String addIcon(Icon image) {
        if (map.containsKey(image)) {
            return map.get(image);
        }

        try {
            File temp = File.createTempFile("temp", Long.toString(System.nanoTime()), tmp);
            ImageIO.write(Utils.createBitmap(image), "png", temp);
            String name = temp.getName();

            map.put(image, name);

            return name;
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

    }

    protected void finalize() throws Throwable {
        super.finalize();

        close();
    }

    public void close() {
        try {
            if (tmp != null) {
                FileUtils.deleteDirectory(tmp);
                tmp = null;
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
