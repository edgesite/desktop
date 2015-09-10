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

Go:

```go
package main

import (
  "github.com/axet/desktop/go"
)

func main() {
  fmt.Println("Home:", desktop.GetHomeFolder())
  fmt.Println("Documents:" desktop.GetDocumentsFolder())
  fmt.Println("AppFolder:" desktop.GetAppDataFolder())
  fmt.Println("Desktop:" desktop.GetDesktopFolder())
  fmt.Println("Downloads:" desktop.GetDownloadsFolder())
}

```
