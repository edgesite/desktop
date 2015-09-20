// +build darwin

package desktop

func desktopMain() {
	app := NSApplicationMainSharedApplication()
	defer app.Release()
	app.Run()
}
