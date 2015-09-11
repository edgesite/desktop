package com.github.axet.desktop;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.image.BufferedImage;
import java.io.IOException;
import java.io.InputStream;

import javax.imageio.ImageIO;
import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JCheckBoxMenuItem;
import javax.swing.JFrame;
import javax.swing.JMenuItem;
import javax.swing.JOptionPane;
import javax.swing.JPopupMenu;
import javax.swing.SwingUtilities;

public class SimpleTrayTest extends JFrame {
    private static final long serialVersionUID = -8634052159132145737L;

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

    public SimpleTrayTest() {
        super("MainFrame");

        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        Icon icon = loadIcon("icon.png");
        Icon icon2 = loadIcon("icon.png");
        Icon icon3 = loadIcon("icon.png");
        Icon icon4 = loadIcon("icon.png");
        Icon icon5 = loadIcon("icon.png");

        menu = new JPopupMenu();
        JMenuItem menuItem1 = new JMenuItem("test disabled");
        menuItem1.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test disabled");
            }
        });
        menuItem1.setEnabled(false);
        menu.add(menuItem1);
        JMenuItem menuItem2 = new JMenuItem("test icon");
        menuItem2.setIcon(icon);
        menuItem2.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test icon");
                SwingUtilities.invokeLater(new Runnable() {
                    @Override
                    public void run() {
                        JOptionPane.showMessageDialog(null, "test icon");
                    }
                });
            }
        });
        menu.add(menuItem2);
        menu.addSeparator();
        final JCheckBoxMenuItem menuItem3 = new JCheckBoxMenuItem("test checkbox");
        menuItem3.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test checkbox");
                sys.update();
            }
        });
        menu.add(menuItem3);
        JMenuItem menuItem4 = new JMenuItem("test normal");
        menuItem4.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test normal");
                SwingUtilities.invokeLater(new Runnable() {
                    @Override
                    public void run() {
                        JOptionPane.showMessageDialog(null, "test normal");
                    }
                });
            }
        });
        menu.add(menuItem4);

        sys.addListener(ml);
        sys.setIcon(icon);
        sys.setTitle("Java tool2");
        sys.setMenu(menu);
        sys.show();

        setSize(300, 200);
        setLocationRelativeTo(null);
        setVisible(true);
    }

    public static void main(String[] args) {
        new SimpleTrayTest();
    }
}
