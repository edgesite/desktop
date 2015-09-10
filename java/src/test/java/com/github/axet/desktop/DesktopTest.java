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
