package com.github.axet.desktop;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JFrame;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

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

    public SimpleTrayTest() {
        super("MainFrame");

        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

        menu = new JPopupMenu();
        JMenuItem menuItem1 = new JMenuItem("test1");
        menuItem1.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test1");
            }
        });
        menu.add(menuItem1);
        JMenuItem menuItem2 = new JMenuItem("test2");
        menuItem2.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test2");
            }
        });
        menu.add(menuItem2);
        menu.addSeparator();
        JMenuItem menuItem3 = new JMenuItem("test3");
        menuItem3.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test3");
            }
        });
        menu.add(menuItem3);
        JMenuItem menuItem4 = new JMenuItem("test4");
        menuItem4.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test4");
            }
        });
        menu.add(menuItem4);

        sys.addListener(ml);
        sys.setIcon(Utils.loadIcon("icon.png"));
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
