package components

import (
	"KiteRunner/internal/model"
	"fmt"

	"github.com/rivo/tview"
)

func ShowModal(a *model.App, title, msg string) {
	m := tview.NewModal().
		SetText(fmt.Sprintf("%s\n\n%s", title, msg)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.Pages.SwitchToPage("login")
			a.Pages.RemovePage("modal")
		})

	a.Pages.AddPage("modal", m, true, true)
}
