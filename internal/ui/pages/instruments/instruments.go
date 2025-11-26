package instruments

import (
	"KiteRunner/internal/model"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Instruments(a *model.App) tview.Primitive {

	// -------------------------
	// 1. Load CSV
	// -------------------------
	instruments, err := LoadInstrumentsCSV("config/mockresponses/instruments_all.csv")
	if err != nil {
		errorBox := tview.NewTextView().
			SetDynamicColors(true).
			SetText("[red]Error loading CSV: " + err.Error())
		return errorBox
	}

	// -------------------------
	// 2. Build Table
	// -------------------------
	table := tview.NewTable().
		SetSelectable(true, false).
		SetBorders(true)

	// HEADER
	headers := []string{
		"Token", "Symbol", "Name", "Price",
		"Strike", "Expiry", "Lot", "Exchange",
	}

	for col, h := range headers {
		table.SetCell(0, col,
			tview.NewTableCell("[yellow::b]"+h).
				SetAlign(tview.AlignCenter),
		)
	}

	// ROWS
	for i, inst := range instruments {
		r := i + 1
		table.SetCell(r, 0, cellInt(inst.InstrumentToken))
		table.SetCell(r, 1, cell(inst.TradingSymbol))
		table.SetCell(r, 2, cell(inst.Name))
		table.SetCell(r, 3, cellFloat(inst.LastPrice))
		table.SetCell(r, 4, cellFloat(inst.Strike))
		table.SetCell(r, 5, cell(inst.Expiry))
		table.SetCell(r, 6, cellInt(inst.LotSize))
		table.SetCell(r, 7, cell(inst.Exchange))
	}

	// Respect NAV / INSERT mode
	table.SetInputCapture(func(ev *tcell.EventKey) *tcell.EventKey {
		if a.Mode == model.ModeInsert {
			return nil // no navigation in Insert mode
		}
		return ev
	})

	return table
}

func cell(text string) *tview.TableCell {
	if text == "NOT AVAILABLE" {
		return tview.NewTableCell("[red]" + text).
			SetAlign(tview.AlignLeft)
	}
	return tview.NewTableCell(text).
		SetAlign(tview.AlignLeft)
}

func cellInt(n int) *tview.TableCell {
	// If 0 but originally missing → turn red
	if n == 0 {
		return tview.NewTableCell("[red]NOT AVAILABLE").
			SetAlign(tview.AlignCenter)
	}
	return tview.NewTableCell(strconv.Itoa(n)).
		SetAlign(tview.AlignCenter)
}

func cellFloat(f float64) *tview.TableCell {
	if f == 0.0 {
		return tview.NewTableCell("[red]NOT AVAILABLE").
			SetAlign(tview.AlignCenter)
	}
	return tview.NewTableCell(fmt.Sprintf("%.2f", f)).
		SetAlign(tview.AlignCenter)
}
