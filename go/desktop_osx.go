// +build !windows

package desktop

import (
  "os"
  
  "./osx/cocoa"
)

// user application data folder
func GetAppDataFolder() string {
  return path(cocoa.NSApplicationSupportDirectory, cocoa.NSUserDomainMask);
}

// user home "/home/user"
func GetHomeFolder() string {
  return os.Getenv("HOME")
}

// user my documents "~/Documents"
func GetDocumentsFolder() string {
  return path(cocoa.NSDocumentDirectory, cocoa.NSUserDomainMask);
}

// user downloads "~/Downloads"
func GetDownloadsFolder() string {
  return path(cocoa.NSDownloadsDirectory, cocoa.NSUserDomainMask)
}

// user desktop "~/Desktop"
func GetDesktopFolder() string {
  return path(cocoa.NSDesktopDirectory, cocoa.NSUserDomainMask);
}

func path(d int , dd int) string {
  f := cocoa.NSFileManagerNew()
  defer f.Release()

  a := f.URLsForDirectoryInDomains(d, dd)
  defer a.Release()

  if a.Count() != 1 {
    return ""
  }

  var u cocoa.NSURL = cocoa.NSURLPointer(a.ObjectAtIndex(0))
  defer u.Release()

  return u.Path()
}
