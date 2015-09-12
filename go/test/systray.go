package main

import (
  "image"
  "time"
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
  
  //menu []desktop.Menu

  s := desktop.DesktopSysTrayNew()
  s.SetIcon(icon)
  s.SetTitle("test")
  //s.SetMenu(menu)
  s.Show()
  
  time.Sleep(5 * time.Second)
}
