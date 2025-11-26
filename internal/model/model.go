package model

import (
	"KiteRunner/internal/config"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIMode int

const (
	ModeNavigation UIMode = iota
	ModeInsert
)

type App struct {
	TUI         *tview.Application
	Pages       *tview.Pages
	UserProfile FullProfile
	Mode        UIMode

	FooterLeft  *tview.TextView
	FooterRight *tview.TextView
}

// Run MUST be in same package as App
func (a *App) Run() error {

	// Set initial footer BEFORE entering event loop
	a.FooterLeft.SetDynamicColors(true)
	a.FooterLeft.SetText("  Mode: [yellow]NAVIGATION  (i = insert, ESC = navigate)").SetTextAlign(tview.AlignLeft)
	a.FooterRight.SetText("  Menu ( m )").SetTextAlign(tview.AlignRight)

	// Deafult satrt mode
	a.Mode = ModeInsert

	a.TUI.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev == nil {
			return ev
		}

		// Toggle into Insert Mode
		if ev.Key() == tcell.KeyRune && ev.Rune() == 'i' && a.Mode == ModeNavigation {
			a.Mode = ModeInsert
			a.UpdateFooter()
			return nil
		}

		// Toggle into Navigation Mode (ESC)
		if ev.Key() == tcell.KeyEsc && a.Mode == ModeInsert {
			a.Mode = ModeNavigation
			a.UpdateFooter()
			return nil
		}

		if ev.Rune() == 'm' && a.Mode == ModeNavigation {
			OpenDashboardMenu(a)
			return nil
		}
		// In Insert mode: let events pass to focused primitive (typing)
		return ev

	})

	// ROOT WRAPPER WITH FOOTER  <--- ADD THIS
	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.Pages, 0, 1, true).
		AddItem(a.FooterLeft, 1, 0, false).
		AddItem(a.FooterRight, 1, 0, false)

	return a.TUI.SetRoot(root, true).Run()
}

func (a *App) UpdateFooter() {
	mode := "[yellow]NAVIGATION"

	if a.Mode == ModeInsert {
		mode = "[green]INSERT"
	}

	a.FooterLeft.SetText(fmt.Sprintf("  Mode: %s   (i = insert, ESC = nav)", mode))
	// Safely redraw only if App is running
	go a.TUI.Draw()
}

type MenuItem struct {
	Title    string
	Shortcut rune
	Action   func(a *App)
}

func BuildMenu(a *App, items []MenuItem) *tview.List {

	// FORCE list to be *tview.List (not Primitive)
	list := tview.NewList()
	list.
		SetSelectedBackgroundColor(tcell.ColorBlue).
		SetBorder(true).
		SetTitle(" Menu ").
		SetTitleAlign(tview.AlignLeft)

	// Add menu items
	for _, item := range items {
		label := fmt.Sprintf("%s (%c)", item.Title, item.Shortcut)

		list.AddItem(label, "", item.Shortcut, func(it MenuItem) func() {
			return func() {
				it.Action(a)
			}
		}(item))
	}

	// Capture ESC to close menu
	list.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev.Key() == tcell.KeyEsc {
			CloseMenu(a)
			return nil
		}
		return ev
	})

	return list
}

func OpenMenu(a *App, menu *tview.List) {
	a.Pages.AddPage("menu", menu, true, true)
}

func CloseMenu(a *App) {
	a.Pages.RemovePage("menu")
}

func OpenDashboardMenu(a *App) {

	items := []MenuItem{
		{
			Title:    "Dashboard",
			Shortcut: rune(config.C.Shortcuts.GotoDashboard[0]),
			Action: func(a *App) {
				a.Pages.SwitchToPage("dashboard")
				CloseMenu(a)
			},
		},
		{
			Title:    "Login Page",
			Shortcut: rune(config.C.Shortcuts.GotoLogin[0]),
			Action: func(a *App) {
				a.Pages.SwitchToPage("login")
				CloseMenu(a)
			},
		},
		{
			Title:    "Close Menu",
			Shortcut: 'x', // add to config if needed
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
				a.Pages.SwitchToPage("instruments")
				CloseMenu(a)
			},
		},
		{
			Title:    "Quotes",
			Shortcut: rune(config.C.Shortcuts.GoToQuotes[0]),
			Action: func(a *App) {
				a.Pages.SwitchToPage("quotes")
				CloseMenu(a)
			},
		},
	}

	menu := BuildMenu(a, items)
	OpenMenu(a, menu)
}
