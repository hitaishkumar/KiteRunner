package orders

import (
	"KiteRunner/internal/model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func OpenOrderTradesModal(a *model.App, trades []model.OrderTrade, order model.Order) {
	// FILTER trades for this specific order
	var filtered []model.OrderTrade
	for _, t := range trades {
		if t.OrderID == order.OrderID {
			filtered = append(filtered, t)
		}
	}

	// BUILD TABLE
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)

	headers := []string{
		"Trade ID", "Txn", "Price", "Qty",
		"Exchange", "Symbol", "Fill TS", "Order TS",
	}

	// HEADER ROW
	for c, h := range headers {
		table.SetCell(0, c,
			tview.NewTableCell("[yellow::b]"+h).
				SetAlign(tview.AlignCenter),
		)
	}

	// ROWS
	for i, t := range filtered {
		r := i + 1
		table.SetCell(r, 0, cell(t.TradeID))
		table.SetCell(r, 1, statusCell(t.TransactionType))
		table.SetCell(r, 2, cellFloat(t.AveragePrice))
		table.SetCell(r, 3, cellInt(t.Quantity))
		table.SetCell(r, 4, cell(t.Exchange))
		table.SetCell(r, 5, cell(t.TradingSymbol))
		table.SetCell(r, 6, cell(t.FillTimestamp))
		table.SetCell(r, 7, cell(t.OrderTimestamp))
	}

	// WRAP IN FRAME (border + title)
	frame := tview.NewFrame(table).
		SetBorders(1, 1, 1, 1, 1, 1)
	frame.SetBorder(true)
	frame.SetTitle(" TRADE BOOK | " + order.OrderID + " | " + order.TradingSymbol + " | ESC/Q to close ")
	frame.SetTitleAlign(tview.AlignLeft)

	// Make the frame modal-sized
	modal := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(frame, 20, 1, true). // height = 20 rows, centered
		AddItem(nil, 0, 1, false)

	// CLOSE HANDLER
	modal.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev.Key() == tcell.KeyEsc || ev.Rune() == 'q' {
			a.Pages.RemovePage("trades_modal")
			return nil
		}
		return ev
	})

	// ADD AS OVERLAY (not a page switch!)
	a.Pages.AddPage("trades_modal", modal, true, true)
	// a.TUI.SetFocus(modal)
}
