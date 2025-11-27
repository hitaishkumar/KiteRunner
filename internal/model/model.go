package model

import (
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
	CurrentPage string

	FooterLeft  *tview.TextView
	FooterRight *tview.TextView
	Footer      *tview.Flex
}

// Run MUST be in same package as App
func (a *App) Run() error {

	// Set initial footer BEFORE entering event loop
	a.CurrentPage = "login"
	a.FooterLeft.SetDynamicColors(true)
	a.FooterRight.SetDynamicColors(true)

	a.Mode = ModeInsert
	a.UpdateFooter()
	a.FooterRight.SetText("[blue]Menu ( m )").SetTextAlign(tview.AlignRight)

	footer := BuildFooter(a.FooterLeft, a.FooterRight)

	// Deafult satrt mode

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

		// In INSERT mode → allow typing
		if a.Mode == ModeInsert {
			return ev
		}

		// In NAV mode → allow only NAV keys
		switch ev.Key() {
		case tcell.KeyUp, tcell.KeyDown, tcell.KeyEnter, tcell.KeyTab, tcell.KeyBacktab:
			return ev
		}

		// In Insert mode: let events pass to focused primitive (typing)
		return ev

	})

	// ROOT WRAPPER WITH FOOTER  <--- ADD THIS
	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.Pages, 0, 1, true).
		AddItem(footer, 1, 0, false)

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

func GlobalMenu(a *App, items []MenuItem) *tview.List {

	// FORCE list to be *tview.List (not Primitive)
	list := tview.NewList()
	list.
		SetSelectedBackgroundColor(tcell.ColorBlue).
		SetBorder(true).
		SetTitle(fmt.Sprintf("%s %s", " Menu ", a.CurrentPage)).
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

	switch a.CurrentPage {
	case "login":
		menu := GlobalMenu(a, GetLoginMenuItems(a))
		OpenMenu(a, menu)
		return
	case "dashboard":
		menu := GlobalMenu(a, GetDashboardMenuItems(a))
		OpenMenu(a, menu)
		return
	case "orders":
		menu := GlobalMenu(a, GetOrdersMenuItems(a))
		OpenMenu(a, menu)
		return
	default:
		menu := GlobalMenu(a, GetOrdersMenuItems(a))
		OpenMenu(a, menu)
		return

	}
}

// BuildFooter creates a footer bar with two text areas: left + right.
func BuildFooter(leftTextView, rightTextView *tview.TextView) tview.Primitive {

	// Flex row with 2 columns
	content := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftTextView, 0, 1, false). // expand left
		AddItem(rightTextView, 0, 1, false) // expand right

	// Overlay content on box using Flex (footerBox is fixed height)
	footer := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(content, 1, 1, false) // height = 1 row
		// If you want border around full footer: add Box behind content
	// footer.SetBorder(true)
	footer.SetBorderColor(tcell.ColorDarkRed)

	return footer
}
