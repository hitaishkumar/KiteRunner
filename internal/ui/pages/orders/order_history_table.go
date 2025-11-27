package orders

import (
	"KiteRunner/internal/model"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildOrderHistoryTableFull(a *model.App, history []model.OrderHistory) *tview.Table {

	table := tview.NewTable().
		SetBorders(true).
		SetFixed(1, 0).
		SetSelectable(true, false)

	headers := []string{
		"Status",
		"Order TS",
		"Order Type",
		"Txn",
		"Price",
		"Trigger",
		"Qty",
		"Filled",
		"Pending",
		"Avg Price",
		"Exchange Order ID",
		"Exchange TS",
		"Product",
		"Variety",
		"Modified",
	}

	// Create header row
	for c, h := range headers {
		table.SetCell(0, c,
			tview.NewTableCell("[yellow::b]"+h).
				SetAlign(tview.AlignCenter),
		)
	}

	// Helper for NA / Zero coloring
	colorize := func(v interface{}) string {
		switch val := v.(type) {
		case string:
			if val == "" {
				return "[red]NOT AVAILABLE"
			}
			return val

		case *string:
			if val == nil || *val == "" {
				return "[red]NOT AVAILABLE"
			}
			return *val

		case int64:
			if val == 0 {
				return "[red]0"
			}
			return fmt.Sprintf("%d", val)

		case float64:
			if val == 0 {
				return "[red]0.00"
			}
			return fmt.Sprintf("%.2f", val)

		default:
			return fmt.Sprintf("%v", val)
		}
	}

	// Fill table rows
	for r, h := range history {
		row := r + 1

		table.SetCell(row, 0, cell(colorize(h.Status)))
		table.SetCell(row, 1, cell(colorize(h.OrderTimestamp)))
		table.SetCell(row, 2, cell(colorize(h.OrderType)))
		table.SetCell(row, 3, cell(colorize(h.TransactionType)))
		table.SetCell(row, 4, cell(colorize(h.Price)))
		table.SetCell(row, 5, cell(colorize(h.TriggerPrice)))
		table.SetCell(row, 6, cell(colorize(h.Quantity)))
		table.SetCell(row, 7, cell(colorize(h.FilledQuantity)))
		table.SetCell(row, 8, cell(colorize(h.PendingQuantity)))
		table.SetCell(row, 9, cell(colorize(h.AveragePrice)))
		table.SetCell(row, 10, cell(colorize(h.ExchangeOrderID)))
		table.SetCell(row, 11, cell(colorize(h.ExchangeTimestamp)))
		table.SetCell(row, 12, cell(colorize(h.Product)))
		table.SetCell(row, 13, cell(colorize(h.Variety)))
		table.SetCell(row, 14, cell(colorize(h.Modified)))
	}

	// Block typing in INSERT mode
	table.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if a.Mode == model.ModeInsert {
			return nil
		}
		return ev
	})

	return table
}

// -------------------------------------------------------------
// ORDER HISTORY MODAL (same style as OrderDetails)
// -------------------------------------------------------------
func OpenOrderHistoryBlockModal(a *model.App, history []model.OrderHistory) {

	// helper for NA / zero fields
	get := func(v interface{}) string {
		switch val := v.(type) {

		case *string:
			if val == nil || *val == "" {
				return "[red]NOT AVAILABLE"
			}
			return *val

		case string:
			if val == "" {
				return "[red]NOT AVAILABLE"
			}
			return val

		case int64:
			if val == 0 {
				return "[red]0"
			}
			return fmt.Sprintf("%d", val)

		case float64:
			if val == 0 {
				return "[red]0.00"
			}
			return fmt.Sprintf("%.2f", val)

		default:
			return fmt.Sprintf("%v", val)
		}
	}

	// ------------------------------------------
	// BUILD RENDERED TEXT (instead of using Flex)
	// ------------------------------------------
	var b strings.Builder

	for idx, h := range history {
		b.WriteString(fmt.Sprintf("[yellow]Entry #%d\n", idx+1))
		b.WriteString("[white]──────────────────────────────────────────────\n")

		b.WriteString(fmt.Sprintf("[yellow]Status: [white]%s\n", h.Status))
		b.WriteString(fmt.Sprintf("[yellow]Timestamp: [white]%s\n", h.OrderTimestamp))
		b.WriteString(fmt.Sprintf("[yellow]Order Type: [white]%s\n", h.OrderType))
		b.WriteString(fmt.Sprintf("[yellow]Txn: [white]%s\n", h.TransactionType))
		b.WriteString(fmt.Sprintf("[yellow]Price: [white]%s\n", get(h.Price)))
		b.WriteString(fmt.Sprintf("[yellow]Trigger Price: [white]%s\n", get(h.TriggerPrice)))
		b.WriteString(fmt.Sprintf("[yellow]Quantity: [white]%s\n", get(h.Quantity)))
		b.WriteString(fmt.Sprintf("[yellow]Filled Qty: [white]%s\n", get(h.FilledQuantity)))
		b.WriteString(fmt.Sprintf("[yellow]Pending Qty: [white]%s\n", get(h.PendingQuantity)))
		b.WriteString(fmt.Sprintf("[yellow]Average Price: [white]%s\n", get(h.AveragePrice)))
		b.WriteString(fmt.Sprintf("[yellow]Exchange Order ID: [white]%s\n", get(h.ExchangeOrderID)))
		b.WriteString(fmt.Sprintf("[yellow]Exchange TS: [white]%s\n", get(h.ExchangeTimestamp)))
		b.WriteString(fmt.Sprintf("[yellow]Product: [white]%s\n", h.Product))
		b.WriteString(fmt.Sprintf("[yellow]Variety: [white]%s\n", h.Variety))
		b.WriteString(fmt.Sprintf("[yellow]Modified: [white]%v\n", h.Modified))

		b.WriteString("[white]──────────────────────────────────────────────\n\n")
	}

	// ------------------------------------------
	// SCROLLABLE TEXT VIEW
	// ------------------------------------------
	scroll := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWrap(true)

	scroll.SetText(b.String())

	scroll.SetChangedFunc(func() { a.TUI.Draw() })

	// ------------------------------------------
	// FRAME AROUND THE TEXT
	// ------------------------------------------
	frame := tview.NewFrame(scroll).
		SetBorders(1, 1, 1, 1, 1, 1)

	frame.SetBorder(true)
	frame.SetTitle(" ORDER HISTORY — Press ESC or Q to Close ")
	frame.SetTitleAlign(tview.AlignLeft)

	frame.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev.Key() == tcell.KeyEsc || ev.Rune() == 'q' {
			a.Pages.RemovePage("order_history")
			return nil
		}
		return ev
	})

	a.Pages.AddAndSwitchToPage("order_history", frame, true)
	a.TUI.SetFocus(scroll)
}

func OpenOrderHistoryTableModal(a *model.App, history []model.OrderHistory, order model.Order) {

	table := buildOrderHistoryTableFull(a, history)

	frame := tview.NewFrame(table).
		SetBorders(1, 1, 1, 1, 1, 1)

	frame.SetBorder(true)

	frame.SetTitle(" ORDER HISTORY (Table View) — ESC/Q to Close | " + order.OrderID + "  |  " + order.TradingSymbol)
	frame.SetTitleAlign(tview.AlignLeft)

	frame.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev.Key() == tcell.KeyEsc || ev.Rune() == 'q' {
			a.Pages.RemovePage("order_history_table")
			return nil
		}
		return ev
	})

	a.Pages.AddAndSwitchToPage("order_history_table", frame, true)
	a.TUI.SetFocus(table)
}
