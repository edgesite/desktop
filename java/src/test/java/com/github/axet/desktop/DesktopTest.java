package com.github.axet.desktop;

public class DesktopTest {
    public static void main(String[] args) {
        // Home folder
        //
        // osx: /Users/user
        // windows: C:\\Users\\user
        // linux: /home/user
        System.out.println("Home: " + Desktop.getHomeFolder());

        // Documents folder
        //
        // osx: /Users/user/Documents
        // windows: C:\\Users\\user\\Documents
        // linux: /home/user/Documents
        System.out.println("Documents: " + Desktop.getDocumentsFolder());

        // Config folder
        //
        // osx: /Users/user/Library/Application Support
        // windows: C:\\Users\\user\\AppData\\Local
        // linux: /home/user/.config
        System.out.println("AppFolder: " + Desktop.getAppDataFolder());

        // Desktop folder
        //
        // osx: /Users/user/Desktop
        // windows: C:\\Users\\user\\Desktop
        // linux: /home/user/Desktop
        System.out.println("Desktop: " + Desktop.getDesktopFolder());

        // Downloads folder
        //
        // osx: /Users/user/Downloads
        // windows: C:\\Users\\user\\Downloads
        // linux: /home/user/Desktop
        System.out.println("Downloads: " + Desktop.getDownloadsFolder());
    }
}
