package orders

import (
	"KiteRunner/internal/model"
	"KiteRunner/internal/ui/components"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// -------------------------------------------------------------
// BUILD ORDERS MAIN TABLE
// -------------------------------------------------------------
func AllOrdersTable(a *model.App, orders []model.Order) *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetFixed(1, 0).
		SetSelectable(true, false)

	// HEADERS
	headers := []string{
		"Seq", "Order ID", "Symbol", "Status",
		"Type", "Txn", "Price", "Qty",
		"Filled", "Pending", "Avg Price",
		"Exchange", "Timestamp",
	}

	for c, h := range headers {
		cell := tview.NewTableCell("[yellow::b]" + h).
			SetAlign(tview.AlignCenter)
		table.SetCell(0, c, cell)
	}

	// ROWS
	for r, o := range orders {
		row := r + 1

		table.SetCell(row, 0, cellInt(r+1))
		table.SetCell(row, 1, cell(o.OrderID))
		table.SetCell(row, 2, cell(o.TradingSymbol))
		table.SetCell(row, 3, statusCell(o.Status))
		table.SetCell(row, 4, cell(o.OrderType))
		table.SetCell(row, 5, statusCell(o.TransactionType)) // BUY/SELL colored
		table.SetCell(row, 6, cellFloat(o.Price))
		table.SetCell(row, 7, cellInt(o.Quantity))
		table.SetCell(row, 8, cellInt(o.FilledQuantity))
		table.SetCell(row, 9, cellInt(o.PendingQuantity))
		table.SetCell(row, 10, cellFloat(o.AveragePrice))
		table.SetCell(row, 11, cell(o.Exchange))
		table.SetCell(row, 12, cell(o.OrderTimestamp))
	}

	// ---------------------------------------------------------
	// INPUT HANDLING (Navigation Mode Only)
	// ---------------------------------------------------------
	table.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {

		// Block all typing in Insert mode
		if a.Mode == model.ModeInsert {
			return nil
		}

		// Handle keys only in NAV mode
		if a.Mode == model.ModeNavigation {

			switch ev.Rune() {

			// ---------------------------------------------------------
			// OPEN ORDER DETAILS CARD (d)
			// ---------------------------------------------------------
			case 'd':
				row, _ := table.GetSelection()
				if row > 0 && row <= len(orders) {
					OpenOrderDetailsModal(a, orders[row-1])
				}
				return nil

			// ---------------------------------------------------------
			// OPEN ORDER HISTORY (h)
			// ---------------------------------------------------------
			case 'h':
				selected, _ := table.GetSelection()
				if selected <= 0 {
					return nil
				}

				// orderID := orders[selected-1].OrderID

				history, err := LoadOrderHistory()
				if err != nil {
					components.ShowModal(a, "Error", "Failed to load order history")
					return nil
				}

				OpenOrderHistoryBlockModal(a, history)
				return nil
			// ---------------------------------------------------------
			// OPEN ORDER HISTORY (h)
			// ---------------------------------------------------------
			case 't':
				selected, _ := table.GetSelection()
				if selected <= 0 {
					return nil
				}

				// orderID := orders[selected-1].OrderID

				history, err := LoadOrderHistory()
				if err != nil {
					components.ShowModal(a, "Error", "Failed to load order history")
					return nil
				}

				row, _ := table.GetSelection()
				if row > 0 && row <= len(orders) {
					OpenOrderHistoryTableModal(a, history, orders[row-1])
				}

				return nil
			case 'o':
				selected, _ := table.GetSelection()
				if selected <= 0 {
					return nil
				}

				// Load trades JSON
				trades, err := LoadOrderTrades() // you already have this loader
				if err != nil {
					components.ShowModal(a, "Error", "Unable to load trades")
					return nil
				}

				OpenOrderTradesModal(a, trades, orders[selected-1])
				return nil

			}
		}

		return ev
	})

	return table
}

// -------------------------------------------------------------
// HELPER CELLS
// -------------------------------------------------------------
func cell(txt string) *tview.TableCell {
	return tview.NewTableCell(txt).SetAlign(tview.AlignCenter)
}

func cellInt(v int) *tview.TableCell {
	return cell(fmt.Sprintf("%d", v))
}

func cellFloat(v float64) *tview.TableCell {
	return cell(fmt.Sprintf("%.2f", v))
}

func statusCell(status string) *tview.TableCell {
	var color string
	switch status {
	case "COMPLETE", "BUY":
		color = "[green]"
	case "REJECTED", "SELL":
		color = "[red]"
	case "CANCELLED":
		color = "[yellow]"
	default:
		color = "[white]"
	}

	return tview.NewTableCell(color + status).
		SetAlign(tview.AlignCenter)
}
