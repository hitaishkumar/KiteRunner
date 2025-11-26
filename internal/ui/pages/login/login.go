package pages

import (
	"KiteRunner/internal/model"
	"KiteRunner/internal/ui/banner"
	"KiteRunner/internal/ui/components"
	"KiteRunner/internal/ui/layout"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Login(a *model.App) tview.Primitive {
	var apiKey, accessToken string

	// ----------------------------------
	// BANNER
	// ----------------------------------
	bannerView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorBlue)
	bannerView.SetText(banner.Banner())

	// ----------------------------------
	// FORM
	// ----------------------------------

	apiField := tview.NewInputField()

	apiField.SetLabel("API Key")
	apiField.SetFieldWidth(40)
	apiField.SetText("")
	apiField.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		// BLOCK input if NOT in insert mode
		if a.Mode == model.ModeNavigation {
			return nil
		}
		return ev
	})

	apiField.SetChangedFunc(func(text string) { apiKey = text })

	accessField := tview.NewInputField()
	accessField.SetLabel("Access Token")
	accessField.SetFieldWidth(40)
	accessField.SetText("")
	accessField.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if a.Mode == model.ModeNavigation {
			return nil
		}
		return ev
	})

	accessField.SetChangedFunc(func(text string) { accessToken = text })

	form := tview.NewForm().
		AddFormItem(apiField).
		AddFormItem(accessField).
		AddButton("Login", func() {

			if apiKey == "" || accessToken == "" {
				components.ShowModal(a, "❌ Login Failed", "API Key or Token is missing")
				return
			}

			a.Pages.SwitchToPage("dashboard")
		}).
		AddButton("Quit", func() { a.TUI.Stop() })

	form.SetBorder(true).
		SetTitle(" Kite Login ").
		SetTitleAlign(tview.AlignLeft)

	// ----------------------------------
	// INPUT MODE TOGGLE LOGIC
	// ----------------------------------

	form.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {

		switch ev.Rune() {
		case 'i': // ENTER INSERT MODE
			a.Mode = model.ModeInsert
			a.UpdateFooter()
			return nil

		case 'n': // ENTER NAVIGATION MODE
			a.Mode = model.ModeNavigation
			a.UpdateFooter()
			return nil
		}

		// In navigation mode block non-nav keys globally
		if a.Mode == model.ModeNavigation {
			// Allow only movement + enter
			switch ev.Key() {
			case tcell.KeyUp, tcell.KeyDown, tcell.KeyEnter, tcell.KeyTab, tcell.KeyBacktab:
				return ev
			}
			return nil
		}

		return ev
	})

	// ----------------------------------
	// FINAL PAGE LAYOUT (banner + form + footer)
	// ----------------------------------
	return layout.WithBanner(bannerView, form, a)
}
