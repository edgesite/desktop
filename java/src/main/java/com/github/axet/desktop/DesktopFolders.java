package com.github.axet.desktop;

import java.io.File;

public interface DesktopFolders {

    /**
         Config folder

         osx: /Users/user/Library/Application Support
         windows: C:\\Users\\user\\AppData\\Local
         linux: /home/user/.config
     */
    abstract public File getAppData();

    /**
         Home folder
        
         osx: /Users/user
         windows: C:\\Users\\user
         linux: /home/user
     */
    abstract public File getHome();

    /**
         Documents folder
        
         osx: /Users/user/Documents
         windows: C:\\Users\\user\\Documents
         linux: /home/user/Documents
     */
    abstract public File getDocuments();

    /**
         Downloads folder
        
         osx: /Users/user/Downloads
         windows: C:\\Users\\user\\Downloads
         linux: /home/user/Desktop
     */
    abstract public File getDownloads();

    /**
         Desktop folder
        
         osx: /Users/user/Desktop
         windows: C:\\Users\\user\\Desktop
         linux: /home/user/Desktop
     */
    abstract public File getDesktop();
}
