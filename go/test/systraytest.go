// +build ignore

package main

import (
	"fmt"
	
	".."
)

const icon_png = `
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
`

const icon2_png =
`
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
`

var icon = desktop.DecodeImageString(icon_png)
var icon2 = desktop.DecodeImageString(icon_png)

type SysTest struct {
	s *desktop.DesktopSysTray
}

func (m *SysTest) Icon1(mn *desktop.Menu) {
	m.s.SetIcon(icon)
	fmt.Println("icon1")
}

func (m *SysTest) Icon2(mn *desktop.Menu) {
	m.s.SetIcon(icon2)
	fmt.Println("icon2")
}

func (m *SysTest) Click(mn *desktop.Menu) {
	fmt.Println("m", mn.Name)
}

func (m *SysTest) Quit(mn *desktop.Menu) {
	fmt.Println("quit", mn.Name)
	m.s.Close()
}

func (m *SysTest) ClickBox(mn *desktop.Menu) {
	fmt.Println(mn.Name)
	mn.State = !mn.State
	m.s.Update()
}

func (m *SysTest) MouseLeftClick() {
	fmt.Println("click")
}

func (m *SysTest) MouseLeftDoubleClick() {
	fmt.Println("dclick")
}

func main() {
	m := &SysTest{desktop.DesktopSysTrayNew()}
	defer m.s.Close()

	menu := []desktop.Menu{
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "Set Icon 1", Action: m.Icon1},
		desktop.Menu{Icon: icon2, Type: desktop.MenuItem, Enabled: true, Name: "Set Icon 2", Action: m.Icon2},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test1", Action: m.Click},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test2", Menu: []desktop.Menu{
			desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test21", Action: m.Click},
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "test22", Action: m.Click},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: false, Name: "test3", Action: m.Click},
		desktop.Menu{Type: desktop.MenuCheckBox, Enabled: true, Name: "test4", State: true, Action: m.ClickBox},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test5", Action: m.Click},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "Quit", Action: m.Quit},
	}

	m.s.AddListener(m)
	m.s.SetIcon(icon)
	m.s.SetTitle("go menu hoho!")
	m.s.SetMenu(menu)
	m.s.Show()

	desktop.Main()
}
