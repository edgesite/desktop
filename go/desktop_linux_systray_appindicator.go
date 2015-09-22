// +build linux

package desktop

import (
	"fmt"
)

type DesktopSysTrayAppIndicator struct {
	*DesktopSysTrayGtk

	App AppIndicator

	IconSet *GtkIconSet

	ShowInvokeVar   GSourceFunc
	HideInvokeVar   GSourceFunc
	UpdateInvokeVar GSourceFunc

	FallbackVar AppIndicatorFallback
}

func DesktopSysTrayAppIndicatorNew(m *DesktopSysTray) *DesktopSysTrayAppIndicator {
	os := &DesktopSysTrayAppIndicator{}
	os.DesktopSysTrayGtk = DesktopSysTrayGtkNew(m)

	os.ShowInvokeVar = func() {
		if os.App == nil {
			os.App = app_indicator_new("SysTrayIcon", "fallback_please", APP_INDICATOR_CATEGORY_APPLICATION_STATUS)
			os.FallbackVar.Set(os.App)
		}

		os.UpdateIcon()

		os.UpdateMenus()
		app_indicator_set_menu(os.App, os.GtkMenu)

		app_indicator_set_status(os.App, APP_INDICATOR_STATUS_ACTIVE)
	}
	os.UpdateInvokeVar = func() {
		os.UpdateIcon()

		os.UpdateMenus()
		app_indicator_set_menu(os.App, os.GtkMenu)

		if os.GtkStatusIcon != nil {
			gtk_status_icon_set_from_gicon(os.GtkStatusIcon, ConvertMenuImage(os.Icon))
			gtk_status_icon_set_tooltip_text(os.GtkStatusIcon, m.Title)
		}
	}
	os.HideInvokeVar = func() {
		if os.GtkStatusIcon != nil {
			gtk_status_icon_set_visible(os.GtkStatusIcon, false)
		}
		app_indicator_set_status(os.App, APP_INDICATOR_STATUS_PASSIVE)
	}

	os.FallbackVar = func() GtkWidget {
		os.GtkStatusIcon = os.CreateGStatusIcon()
		gtk_status_icon_set_visible(os.GtkStatusIcon, true)
		return os.GtkStatusIcon
	}

	return os
}

func (os *DesktopSysTrayAppIndicator) show() {
	GtkMessageLoopInvoke(&os.ShowInvokeVar)
}

func (os *DesktopSysTrayAppIndicator) hide() {
	GtkMessageLoopInvoke(&os.HideInvokeVar)
}

func (os *DesktopSysTrayAppIndicator) update() {
	GtkMessageLoopInvoke(&os.UpdateInvokeVar)
}

func (os *DesktopSysTrayAppIndicator) close() {
	os.FallbackVar.Close(os.App)
	os.IconSet.Close()
	os.DesktopSysTrayGtk.close()
}

func (os *DesktopSysTrayAppIndicator) UpdateIcon() {
	if os.Icon == nil {
		return
	}

	if os.IconSet == nil {
		os.IconSet = GtkIconSetNew()
	}

	p := os.IconSet.Add(os.Icon)
	fmt.Println(os.IconSet.Path, p)
	app_indicator_set_icon_theme_path(os.App, os.IconSet.Path)
	app_indicator_set_icon_full(os.App, p, "SysTrayIcon")
}

