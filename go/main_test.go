package desktop

import (
  "testing"
  "fmt"
)

func Test(t *testing.T) {
  fmt.Println("Home:", GetHomeFolder())
  fmt.Println("Documents:", GetDocumentsFolder())
  fmt.Println("AppFolder:", GetAppDataFolder())
  fmt.Println("Desktop:", GetDesktopFolder())
  fmt.Println("Downloads:", GetDownloadsFolder())
}
