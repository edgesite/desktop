package main

import (
	"image"
	"os"

	".."
)

func main() {
	file, err := os.Open("icon.png")
	if err != nil {
		panic(err)
	}
	icon, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	menu := []desktop.Menu{
    desktop.Menu{Type:desktop.MenuItem, Enabled:true, Name:"test1"},
    desktop.Menu{Type:desktop.MenuItem, Enabled:true, Name:"test2"},
	}

	s := desktop.DesktopSysTrayNew()
	s.SetIcon(icon)
	s.SetTitle("test")
	s.SetMenu(menu)
	s.Show()
  
  desktop.Main()
}
