package com.github.axet.desktop;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.image.BufferedImage;
import java.io.IOException;
import java.io.InputStream;

import javax.imageio.ImageIO;
import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

public class SimpleTrayTest {
    DesktopSysTray sys = Desktop.getDesktopSysTray();
    JPopupMenu menu;

    DesktopSysTray.Listener ml = new DesktopSysTray.Listener() {
        @Override
        public void mouseLeftClick() {
            System.out.println("left click");
        }

        @Override
        public void mouseLeftDoubleClick() {
            System.out.println("double click");
        }
    };

    ImageIcon loadImage(String s) {
        try {
            InputStream is = getClass().getResourceAsStream(s);
            BufferedImage bmp = ImageIO.read(is);
            return new ImageIcon(bmp);
        } catch (IOException e1) {
            throw new RuntimeException(e1);
        }
    }

    public SimpleTrayTest() {
        final Icon icon1 = loadImage("icon.png");
        final Icon icon2 = loadImage("icon2.png");

        menu = new JPopupMenu();
        JMenuItem menuItem1 = new JMenuItem("icon1");
        menuItem1.setIcon(icon1);
        menuItem1.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                sys.setIcon(icon1);
                System.out.println("icon1");
            }
        });
        menu.add(menuItem1);
        menu.addSeparator();
        
        JMenuItem menuItem2 = new JMenuItem("icon2");
        menuItem2.setIcon(icon2);
        menuItem2.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                sys.setIcon(icon2);
                System.out.println("icon2");
            }
        });
        menu.add(menuItem2);

        JMenuItem menuItem3 = new JMenuItem("Quit");
        menuItem3.setIcon(icon2);
        menuItem3.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                sys.close();
                System.out.println("Quit");
            }
        });
        menu.add(menuItem3);

        sys.addListener(ml);
        sys.setTitle("Java tool2");
        sys.setIcon(icon1);
        sys.setMenu(menu);
        sys.show();
    }

    public static void main(String[] args) {
        new SimpleTrayTest();
    }
}