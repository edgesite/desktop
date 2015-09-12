# Desktop

Java desktop functions. Have you tried to find user default Download folder using java? If so, you would find this
library very helpful.

## Example Desktop Folders

```java    
package com.github.axet.desktop;

public class DesktopTest {
    public static void main(String[] args) {
        DesktopFolders d = Desktop.getDesktopFolders();

        // Home folder: /Users/user
        System.out.println("Home: " + d.getHome());
        // Documents folder /Users/user/Documents
        System.out.println("Documents: " + d.getDocuments());
        // Config folder /Users/axet/Library/Application Support
        System.out.println("AppFolder: " + d.getAppData());
        // Desktop folder /Users/axet/Desktop
        System.out.println("Desktop: " + d.getDesktop());
        // Downloads folder /Users/axet/Downloads
        System.out.println("Downloads: " + d.getDownloads());
    }
}
```

## Example Sys Tray Icon
(aka Notification Area Icons or Status Bar icons)

Supports: JMenu (submenu), JCheckBoxMenuItem, JMenuItem, JPopupMenu.Separator, setImage() and setEnabled() on JMenuItem.

Full example with submenus, icons, checkboxes can be found in test folder:

  * [SysTrayTest.java](src/test/java/com/github/axet/desktop/SysTrayTest.java)

```java
package com.github.axet.desktop;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.image.BufferedImage;
import java.io.IOException;
import java.io.InputStream;
import java.util.List;

import javax.imageio.ImageIO;
import javax.swing.Icon;
import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

import net.sf.image4j.codec.ico.ICODecoder;

public class SimpleTrayTest extends JFrame {
    private static final long serialVersionUID = 1L;

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
        menu.addSeparator();
        JMenuItem menuItem2 = new JMenuItem("test2");
        menuItem2.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent arg0) {
                System.out.println("test2");
            }
        });
        menu.add(menuItem2);

        Icon icon = null;

        try {
            InputStream is = getClass().getResourceAsStream("bug.ico");
            List<BufferedImage> bmp = ICODecoder.read(is);
            icon = new ImageIcon(bmp.get(0));
        } catch (IOException e1) {
            throw new RuntimeException(e1);
        }

        sys.addListener(ml);
        sys.setTitle("Java tool2");
        sys.setIcon(icon);
        sys.setMenu(menu);
        sys.show();
    }

    public static void main(String[] args) {
        new SimpleTrayTest();
    }
}
```

## Central Maven Repo

```xml
<dependencies>
	<dependency>
	  <groupId>com.github.axet</groupId>
	  <artifactId>desktop</artifactId>
	  <version>2.2.5</version>
	</dependency>
</dependencies>
```