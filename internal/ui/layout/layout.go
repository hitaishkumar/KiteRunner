package layout

import (
	"KiteRunner/internal/model"
	"fmt"

	"github.com/rivo/tview"
)

func WithBanner(banner tview.Primitive, content tview.Primitive, app *model.App) *tview.Flex {

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(banner, 10, 1, false).
		AddItem(content, 0, 2, true)
}

func WithContextBanner(banner tview.Primitive, content tview.Primitive, app *model.App) *tview.Flex {
	// ----------------------------------
	// USER PANEL
	// ----------------------------------
	userPanel := tview.NewTextView().SetDynamicColors(true)

	userPanel.SetText(
		fmt.Sprintf("[blue]User ID:[white] %s\n", app.UserProfile.Data.UserID) +
			fmt.Sprintf("[blue]Name:   [white] %s\n", app.UserProfile.Data.UserName) +
			fmt.Sprintf("[blue]Email:  [white] %s\n", app.UserProfile.Data.Email) +
			fmt.Sprintf("[blue]Phone:  [white] %s\n", app.UserProfile.Data.Phone) +
			fmt.Sprintf("[blue]PAN:    [white] %s\n", app.UserProfile.Data.Pan) +
			fmt.Sprintf("[blue]Broker: [white] %s", app.UserProfile.Data.Broker),
	)
	userPanel.SetBorder(true).SetTitle("User Info")

	// ----------------------------------
	// TOP ROW = Banner + User Panel (SIDE BY SIDE)
	// ----------------------------------
	topBar := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(banner, 100, 0, false). // banner width
		AddItem(userPanel, 0, 1, false) // user panel takes remaining width

	// ----------------------------------
	// FINAL LAYOUT (TopRow + Content)
	// ----------------------------------
	contenRow := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(topBar, 8, 0, false). // height of the top section
		AddItem(content, 0, 1, true)  // content fills rest vertically

	return contenRow
}
