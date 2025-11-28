package orders

import (
	"KiteRunner/internal/model"

	"github.com/rivo/tview"
)

func Page(a *model.App) tview.Primitive {

	orders, err := LoadOrders()
	if err != nil {
		errView := tview.NewTextView().SetDynamicColors(true)
		errView.SetText("[red]Failed to load orders: " + err.Error())
		return errView
	}

	table := AllOrdersTable(a, orders)

	page := tview.NewFlex()
	page.SetDirection(tview.FlexRow).
		SetTitle("Order Book"+" | "+a.UserProfile.Data.UserName+"  |  "+"Press h (history) | d (more detail) | t (order history) | o (order's trades)  ").SetBorder(true).SetTitleAlign(tview.AlignLeft).SetBorderPadding(2, 2, 4, 0)
	// AddItem(banner, 10, 1, false).
	page.AddItem(table, 0, 2, true)

	return page

}
