package com.github.axet.desktop.os.linux.handle;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.HashMap;

import javax.imageio.ImageIO;
import javax.swing.Icon;

import org.apache.commons.io.FileUtils;
import org.apache.commons.lang.StringUtils;

import com.github.axet.desktop.Utils;

public class GtkIconSet {
    File tmp;
    HashMap<Icon, String> map = new HashMap<Icon, String>();

    public GtkIconSet() {
        try {
            tmp = Files.createTempDirectory("systrayicon").toFile();
            tmp.deleteOnExit();
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
            String format = "png";
            String suffix = "." + format;

            File temp = File.createTempFile("temp", suffix, tmp);
            temp.deleteOnExit();

            ImageIO.write(Utils.createBitmap(image), format, temp);
            String name = temp.getName();

            name = StringUtils.removeEnd(name, suffix);

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
