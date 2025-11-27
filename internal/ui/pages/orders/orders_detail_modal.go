package orders

import (
	"KiteRunner/internal/model"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// -------------------------------------------------------------
// ORDER DETAILS MODAL (card)
// -------------------------------------------------------------
func OpenOrderDetailsModal(a *model.App, o model.Order) {

	// A reusable function for nil & empty handling + red color for NA/zero
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

		case bool:
			return fmt.Sprintf("%v", val)

		case []string:
			if len(val) == 0 {
				return "[red]NOT AVAILABLE"
			}
			return strings.Join(val, ", ")

		case map[string]interface{}:
			if len(val) == 0 {
				return "[red]NOT AVAILABLE"
			}
			b, _ := json.MarshalIndent(val, "", "  ")
			return string(b)

		default:
			return fmt.Sprintf("%v", val)
		}
	}

	// -------------------------
	// BUILD CONTENT FLEX
	// -------------------------
	content := tview.NewFlex().SetDirection(tview.FlexRow)

	addRow := func(label, value string) {

		// ------------ ROW ------------
		row := tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(
				tview.NewTextView().
					SetDynamicColors(true).
					SetText("[yellow]"+label),
				20, 1, false,
			).
			AddItem(
				tview.NewTextView().
					SetDynamicColors(true).
					SetText(value),
				0, 3, false,
			)

		// ------------ SEPARATOR ------------
		// sep := sxt("[grey]────────────────────────────────────────────────────────────────────────")

		// Height 1 ⇒ looks like a bottom border
		content.AddItem(row, 1, 0, false)
		// content.AddItem(sep, 1, 0, false)
	}

	// ----------- ALL FIELDS -----------
	addRow("Placed By", o.PlacedBy)
	addRow("Order ID", o.OrderID)
	addRow("Exchange Order ID", get(o.ExchangeOrderID))
	addRow("Parent Order ID", get(o.ParentOrderID))
	addRow("Status", o.Status)
	addRow("Status Message", get(o.StatusMessage))
	addRow("Status Msg Raw", get(o.StatusMessageRaw))
	addRow("Order Timestamp", o.OrderTimestamp)
	addRow("Exchange Update TS", get(o.ExchangeUpdateTime))
	addRow("Exchange TS", get(o.ExchangeTimestamp))
	addRow("Variety", o.Variety)
	addRow("Modified", fmt.Sprintf("%v", o.Modified))
	addRow("Exchange", o.Exchange)
	addRow("Trading Symbol", o.TradingSymbol)
	addRow("Instrument Token", get(o.InstrumentToken))
	addRow("Order Type", o.OrderType)
	addRow("Transaction Type", o.TransactionType)
	addRow("Validity", o.Validity)
	addRow("Validity TTL", get(o.ValidityTTL))
	addRow("Product", o.Product)
	addRow("Quantity", get(o.Quantity))
	addRow("Disclosed Quantity", get(o.DisclosedQuantity))
	addRow("Price", get(o.Price))
	addRow("Trigger Price", get(o.TriggerPrice))
	addRow("Average Price", get(o.AveragePrice))
	addRow("Filled Qty", get(o.FilledQuantity))
	addRow("Pending Qty", get(o.PendingQuantity))
	addRow("Cancelled Qty", get(o.CancelledQuantity))
	addRow("Auction Number", get(o.AuctionNumber))
	addRow("Market Protection", get(o.MarketProtection))
	addRow("Tag", get(o.Tag))
	addRow("GUID", get(o.GUID))
	addRow("Tags", get(o.Tags))
	addRow("Meta", get(o.Meta))

	// -------------------------
	// NOW WRAP WITH FRAME (border + title)
	// -------------------------
	frame := tview.NewFrame(content).
		SetBorders(1, 1, 1, 1, 1, 1)

	frame.SetBorder(true)
	frame.SetTitle(" ORDER DETAILS | " + o.OrderID + " | " + o.TradingSymbol + " | Press ESC or Q to close ")
	frame.SetTitleAlign(tview.AlignLeft)

	// Key handling to close modal
	frame.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if ev.Key() == tcell.KeyEsc || ev.Rune() == 'q' {
			a.Pages.RemovePage("order_details")
			return nil
		}
		return ev
	})

	// Show modal page
	a.Pages.AddAndSwitchToPage("order_details", frame, true)
	a.TUI.SetFocus(frame)
}
