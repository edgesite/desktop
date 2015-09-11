// +build darwin

package desktop

import (
  "os"
)

// user application data folder
func GetAppDataFolder() string {
  return path(NSApplicationSupportDirectory, NSUserDomainMask);
}

// user home "/home/user"
func GetHomeFolder() string {
  return os.Getenv("HOME")
}

// user my documents "~/Documents"
func GetDocumentsFolder() string {
  return path(NSDocumentDirectory, NSUserDomainMask);
}

// user downloads "~/Downloads"
func GetDownloadsFolder() string {
  return path(NSDownloadsDirectory, NSUserDomainMask)
}

// user desktop "~/Desktop"
func GetDesktopFolder() string {
  return path(NSDesktopDirectory, NSUserDomainMask);
}

func path(d int , dd int) string {
  f := NSFileManagerNew()
  defer f.Release()

  a := f.URLsForDirectoryInDomains(d, dd)
  defer a.Release()

  if a.Count() != 1 {
    return ""
  }

  var u NSURL = NSURLPointer(a.ObjectAtIndex(0))
  defer u.Release()

  return u.Path()
}
