package streaming

import (
	"KiteRunner/internal/model"

	"github.com/rivo/tview"
)

// ---- TVIEW WRAPPER PAGE ----

func StreamingChartsPage(a *model.App) tview.Primitive {

	textView := tview.NewTextView()

	textView.SetDynamicColors(true).
		SetBorder(true)

	textView.SetTitle(" Live Chart ")
	textView.SetTitleAlign(tview.AlignLeft)

	return textView
}
