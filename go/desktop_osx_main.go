// +build darwin

package desktop

var App NSApplication

func desktopMain() {
	App = NSApplicationMainSharedApplication()
	defer App.Release()
	App.Run()
}

func desktopMainClose() {
  App.Terminate()
}
