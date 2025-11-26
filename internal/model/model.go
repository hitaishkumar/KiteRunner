package model

import (
	"fmt"
	"strings"

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

func BuildFuzzyModal(a *App, items []MenuItem) *tview.Frame {
	inputField := tview.NewInputField()
	inputField.SetBorder(true).SetBorderPadding(0, 0, 1, 1)
	inputField.SetFieldBackgroundColor(tcell.ColorBlack)

	list := tview.NewList()
	list.SetBorder(true).SetBorderPadding(1, 1, 1, 1)

	// maintain filteredItems state to invoke action on ENTER
	filteredItems := make([]MenuItem, 0)
	filterHandler := func(currentText string) {
		list.Clear()
		filteredItems = filteredItems[:0]
		for _, item := range items {
			if strings.Contains(strings.ToLower(item.Title), currentText) {
				filteredItems = append(filteredItems, item)
				list.AddItem(item.Title, "", 0, nil)
			}
		}
	}
	inputField.SetChangedFunc(filterHandler)
	// initalize filter as empty
	filterHandler("")

	// hanlde forward input enter to focused items' action invokation
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			idx := list.GetCurrentItem()
			if idx >= 0 {
				selectedItem := filteredItems[idx]
				if selectedItem.Action != nil {
					// invoke the action
					selectedItem.Action(a)
					// close the modal
					CloseFuzzyModal(a)
				}
			}
		}
	})

	modalContent := tview.NewFlex()
	modalContent.SetDirection(tview.FlexRow)
	modalContent.AddItem(inputField, 3, 0, true)
	modalContent.AddItem(list, 0, 1, false)

	frame := tview.NewFrame(modalContent)
	frame.SetBorders(1, 1, 1, 1, 2, 2)
	frame.SetBorder(true)
	frame.SetTitle(fmt.Sprintf("%s %s", a.CurrentPage, "Menu"))
	frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			CloseFuzzyModal(a)
			return nil
		}
		return event
	})

	return frame
}

func OpenFuzzyModal(a *App, modal *tview.Frame) {
	centered := tview.NewGrid().
		SetRows(0, 30, 0).    // middle row height
		SetColumns(0, 80, 0). // middle column width
		AddItem(modal, 1, 1, 1, 1, 0, 0, true)
	a.Pages.AddPage("fuzzy menu", centered, true, true)
}

func CloseFuzzyModal(a *App) {
	a.Pages.RemovePage("fuzzy menu")
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
		menu := BuildFuzzyModal(a, GetLoginMenuItems(a))
		OpenFuzzyModal(a, menu)
		return
	case "dashboard":
		menu := BuildFuzzyModal(a, GetDashboardMenuItems(a))
		OpenFuzzyModal(a, menu)
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
