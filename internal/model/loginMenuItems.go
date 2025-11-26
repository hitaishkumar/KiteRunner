package model

import "KiteRunner/internal/config"

func GetLoginMenuItems(a *App) []MenuItem {
	return []MenuItem{
		{
			Title:    "Dashboard",
			Shortcut: rune(config.C.Shortcuts.GotoDashboard[0]),
			Action: func(a *App) {
				a.CurrentPage = "dashboard"
				a.Pages.SwitchToPage("dashboard")
				CloseMenu(a)
			},
		},
		{
			Title:    "Login Page",
			Shortcut: rune(config.C.Shortcuts.GotoLogin[0]),
			Action: func(a *App) {
				a.CurrentPage = "login"
				a.Pages.SwitchToPage("login")
				CloseMenu(a)
			},
		},
		{
			Title:    "Close Menu",
			Shortcut: 'x',
			Action: func(a *App) {
				CloseMenu(a)
			},
		},
		{
			Title:    "Quit",
			Shortcut: rune(config.C.Shortcuts.Quit[0]),
			Action: func(a *App) {
				a.TUI.Stop()
			},
		},
		{
			Title:    "Instruments",
			Shortcut: rune(config.C.Shortcuts.GoToInstruments[0]),
			Action: func(a *App) {
				a.CurrentPage = "instruments"
				a.Pages.SwitchToPage("instruments")
				CloseMenu(a)
			},
		},
		{
			Title:    "Quotes",
			Shortcut: rune(config.C.Shortcuts.GoToQuotes[0]),
			Action: func(a *App) {
				a.CurrentPage = "quotes"
				a.Pages.SwitchToPage("quotes")
				CloseMenu(a)
			},
		},
	}
}

func GetDashboardMenuItems(a *App) []MenuItem {
	return []MenuItem{
		{
			Title:    "GO to Login",
			Shortcut: rune(config.C.Shortcuts.GotoDashboard[0]),
			Action: func(a *App) {
				a.CurrentPage = "login"
				a.Pages.SwitchToPage("login")
				CloseMenu(a)
			},
		},
	}
}
