# Desktop

Java && Go Desktop functions. Have you tried to find user default Download folder using java? It is very complicated. If so, you would find this library very helpful.

## Features

  * Cross platform user specific folders: Download folder, Home folder, Documents folder
  * Cross platform SysTray Icon
  * Cross platform Browser Pop Window

## Example Desktop Folders

Java:

```java
package com.github.axet.desktop;

public class DesktopTest {
    public static void main(String[] args) {
        DesktopFolders d = Desktop.getDesktopFolders();

        // Home folder
        //
        // osx: /Users/user
        // windows: C:\Users\user
        // linux: /home/user
        System.out.println("Home: " + d.getHome());
        
        // Documents folder
        //
        // osx: /Users/user/Documents
        // windows: C:\Users\user\Documents
        // linux: /home/user/Documents
        System.out.println("Documents: " + d.getDocuments());

        // Config folder
        //
        // osx: /Users/user/Library/Application Support
        // windows: C:\Users\user\AppData\Local
        // linux: /home/user/.config
        System.out.println("AppFolder: " + d.getAppData());

        // Desktop folder
        //
        // osx: /Users/user/Desktop
        // windows: C:\Users\user\Desktop
        // linux: /home/user/Desktop
        System.out.println("Desktop: " + d.getDesktop());
        
        // Downloads folder
        //
        // osx: /Users/user/Downloads
        // windows: C:\Users\user\Downloads
        // linux: /home/user/Desktop
        System.out.println("Downloads: " + d.getDownloads());
    }
}
```

Go:

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
