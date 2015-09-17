// +build ignore

package main

import (
	".."
	"fmt"
)

func main() {
	fmt.Println("Home:", desktop.GetHomeFolder())
	fmt.Println("Documents:", desktop.GetDocumentsFolder())
	fmt.Println("AppFolder:", desktop.GetAppDataFolder())
	fmt.Println("Desktop:", desktop.GetDesktopFolder())
	fmt.Println("Downloads:", desktop.GetDownloadsFolder())
}
