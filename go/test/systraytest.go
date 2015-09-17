// +build ignore

package main

import (
	"fmt"
	
	".."
)

const icon_png = `
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
`

type SysTest struct {
	s *desktop.DesktopSysTray
}

func (m *SysTest) Click(mn *desktop.Menu) {
	fmt.Println("m", mn.Name)
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

	icon := desktop.DecodeImageString(icon_png)

	menu := []desktop.Menu{
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test1", Action: m.Click},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test2", Menu: []desktop.Menu{
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "test21", Action: m.Click},
			desktop.Menu{Type: desktop.MenuItem, Enabled: true, Name: "test22", Action: m.Click},
		}},
		desktop.Menu{Type: desktop.MenuItem, Enabled: false, Name: "test3", Action: m.Click},
		desktop.Menu{Type: desktop.MenuCheckBox, Enabled: true, Name: "test4", State: true, Action: m.ClickBox},
		desktop.Menu{Type: desktop.MenuSeparator},
		desktop.Menu{Icon: icon, Type: desktop.MenuItem, Enabled: true, Name: "test5", Action: m.Click},
	}

	m.s.AddListener(m)
	m.s.SetIcon(icon)
	m.s.SetTitle("go menu hoho!")
	m.s.SetMenu(menu)
	m.s.Show()

	desktop.Main()
}
