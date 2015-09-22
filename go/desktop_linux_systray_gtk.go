// +build linux

package desktop

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/png"
)

var SpaceIcon image.Image = image.NewRGBA(image.Rect(0, 0, 1, 1))

type DesktopSysTrayGtk struct {
	M             *DesktopSysTray
	Icon          image.Image
	GtkStatusIcon GtkWidget
	GtkMenu       GtkWidget

	IconActivate GSourceFunc
	IconPopup    GSourceFunc

	ShowInvokeVar   GSourceFunc
	HideInvokeVar   GSourceFunc
	UpdateInvokeVar GSourceFunc

	GSourceFuncs []*GSourceFunc
}

func DesktopSysTrayGtkNew(m *DesktopSysTray) *DesktopSysTrayGtk {
	GtkMessageLoopInc()

	os := &DesktopSysTrayGtk{}

	os.M = m

	os.ShowInvokeVar = func() {
		os.UpdateMenus()
		if os.GtkStatusIcon == nil {
			os.GtkStatusIcon = os.CreateGStatusIcon()
		}
		gtk_status_icon_set_visible(os.GtkStatusIcon, true)
	}

	os.HideInvokeVar = func() {
		gtk_status_icon_set_visible(os.GtkStatusIcon, false)
	}

	os.UpdateInvokeVar = func() {
		os.UpdateMenus()

		if os.GtkStatusIcon != nil {
			gtk_status_icon_set_from_gicon(os.GtkStatusIcon, ConvertMenuImage(os.Icon))
			gtk_status_icon_set_tooltip_text(os.GtkStatusIcon, m.Title)
		}
	}

	os.IconActivate = func() {
		for l := range m.Listeners {
			l.MouseLeftClick()
		}
	}

	os.IconPopup = func() {
		gtk_menu_popup(os.GtkMenu, nil, nil, gtk_status_icon_position_menu, nil, 1, gtk_get_current_event_time())
	}

	return os
}

func ConvertMenuImage(icon image.Image) GIcon {
	var menubarHeigh uint = 64

	c := resize.Resize(menubarHeigh, menubarHeigh, icon, resize.Lanczos3)

	var b bytes.Buffer
	err := png.Encode(&b, c)
	if err != nil {
		panic(err)
	}

	buf := b.Bytes()
	gb := g_bytes_new(buf, len(buf))
	gi := g_bytes_icon_new(gb)
	return gi
}

func (os *DesktopSysTrayGtk) CreateMenuItem(item *Menu) GtkWidget {
	img := item.Icon

	if img == nil {
		img = SpaceIcon
	}

	spacing := 6

	box := gtk_hbox_new(false, spacing)
	wicon := gtk_image_new_from_gicon(ConvertMenuImage(img), GTK_ICON_SIZE_MENU)
	gtk_box_pack_start(box, wicon, false, false, spacing)
	label := gtk_label_new(item.Name)

	var menu GtkWidget

	switch item.Type {
	case MenuCheckBox:
		menu = gtk_check_menu_item_new()
		gtk_check_menu_item_set_active(menu, item.State)
	case MenuItem:
		menu = gtk_menu_item_new()
	}

	if !item.Enabled {
		gtk_widget_set_sensitive(menu, false)
	}

	gtk_box_pack_start(box, label, false, false, spacing)
	gtk_container_add(menu, box)
	gtk_widget_show_all(menu)

	if item.Menu == nil {
		var fn GSourceFunc = func() {
			item.Action(item)
		}
		os.GSourceFuncs = append(os.GSourceFuncs, &fn)
		g_signal_connect_activate(menu, &fn)
	}

	return menu
}

func (os *DesktopSysTrayGtk) CreateSubMenu(mm []Menu) GtkWidget {
	gmenu := gtk_menu_new()

	for i := range mm {
		mn := &mm[i]

		switch mn.Type {
		case MenuItem, MenuCheckBox:
			if mn.Menu != nil {
				sub := os.CreateSubMenu(mn.Menu)
				item := os.CreateMenuItem(mn)
				gtk_menu_item_set_submenu(item, sub)
				gtk_menu_shell_append(gmenu, item)
			} else {
				item := os.CreateMenuItem(mn)
				gtk_menu_shell_append(gmenu, item)
			}
		case MenuSeparator:
			item := gtk_separator_menu_item_new()
			gtk_menu_shell_append(gmenu, item)
		}
	}

	return gmenu
}

func (os *DesktopSysTrayGtk) UpdateMenus() {
	m := os.M

	if os.GtkMenu != nil {
		gtk_widget_destroy(os.GtkMenu)
	}

	os.GSourceFuncs = nil

	os.GtkMenu = os.CreateSubMenu(m.Menu)
}

func (os *DesktopSysTrayGtk) CreateGStatusIcon() GtkWidget {
	m := os.M

	gicon := gtk_status_icon_new_from_gicon(ConvertMenuImage(os.Icon))

	g_signal_connect_activate(gicon, &os.IconActivate)
	g_signal_connect_popup(gicon, &os.IconPopup)

	gtk_status_icon_set_tooltip_text(gicon, m.Title)
	return gicon
}

func (os *DesktopSysTrayGtk) show() {
	GtkMessageLoopInvoke(&os.ShowInvokeVar)
}

func (os *DesktopSysTrayGtk) hide() {
	GtkMessageLoopInvoke(&os.HideInvokeVar)
}

func (os *DesktopSysTrayGtk) update() {
	GtkMessageLoopInvoke(&os.UpdateInvokeVar)
}

func (os *DesktopSysTrayGtk) close() {
	if os.GtkMenu != nil {
		gtk_widget_destroy(os.GtkMenu)
		os.GtkMenu = nil
	}

	if os.GtkStatusIcon != nil {
		gtk_widget_destroy(os.GtkStatusIcon)
		os.GtkStatusIcon = nil
	}

	os.GSourceFuncs = nil

	GtkMessageLoopDec()
}

func (os *DesktopSysTrayGtk) setIcon(i image.Image) {
	os.Icon = i
}

