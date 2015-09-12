# Desktop

Java desktop functions. Have you tried to find user default Download folder using java? If so, you would find this
library very helpful.

## Example Desktop Folders
    
```go
package main

import (
  "github.com/axet/desktop/go"
)

func main() {
  // Home folder
  //
  // osx: /Users/user
  // windows: C:\Users\user
  // linux: /home/user
  fmt.Println("Home:", desktop.GetHomeFolder())
  
  // Documents folder
  //
  // osx: /Users/user/Documents
  // windows: C:\Users\user\Documents
  // linux: /home/user/Documents
  fmt.Println("Documents:" desktop.GetDocumentsFolder())
  
  // Config folder
  //
  // osx: /Users/user/Library/Application Support
  // windows: C:\Users\user\AppData\Local
  // linux: /home/user/.config
  fmt.Println("AppFolder:" desktop.GetAppDataFolder())
  
  // Desktop folder
  //
  // osx: /Users/user/Desktop
  // windows: C:\Users\user\Desktop
  // linux: /home/user/Desktop
  fmt.Println("Desktop:" desktop.GetDesktopFolder())
  
  // Downloads folder
  //
  // osx: /Users/user/Downloads
  // windows: C:\Users\user\Downloads
  // linux: /home/user/Desktop
  fmt.Println("Downloads:" desktop.GetDownloadsFolder())
}
```

## Example Sys Tray Icon
(aka Notification Area Icons or Status Bar icons)

```java
package com.github.axet.desktop;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JFrame;
import javax.swing.JMenuItem;
import javax.swing.JPopupMenu;

public class SysTrayTest extends JFrame {

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

        @Override
        public void mouseRightClick() {
            System.out.println("right click");
            sys.showContextMenu();
        }

    };

    public SysTrayTest() {
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
        menu.addSeparator();

        sys.addListener(ml);
        sys.setTitle("Java tool2");
        sys.setMenu(menu);
        sys.show();
    }

    public static void main(String[] args) {
        new SysTrayTest();
    }
}
```