package quotes

import (
	"fmt"

	"KiteRunner/internal/model"

	"github.com/rivo/tview"
)

func Quotes(a *model.App) tview.Primitive {

	quoteData, err := LoadQuote("config/mockresponses/quote.json")
	if err != nil {
		return tview.NewTextView().SetText("[red]Failed to load quote.json: " + err.Error())
	}

	// get first symbol
	var symbol string
	var data model.InstrumentQuote
	for k, v := range quoteData.Data {
		symbol = k
		data = v
		break
	}

	// ======================================
	// LTP PANEL
	// ======================================
	ltpPanel := tview.NewTextView()
	ltpPanel.SetDynamicColors(true)
	ltpPanel.SetBorder(true)
	ltpPanel.SetTitle(" Quote: " + symbol)

	ltpPanel.SetText(fmt.Sprintf(
		"[yellow]Last Price:[white] %.2f\n"+
			"[green]Open:[white] %.2f  [green]High:[white] %.2f\n"+
			"[red]Low:[white] %.2f   [red]Close:[white] %.2f\n\n"+
			"[yellow]Volume:[white] %d\n"+
			"[yellow]OI:[white] %d\n"+
			"[yellow]Net Change:[white] %.2f\n\n"+
			"[yellow]Lower Circuit:[white] %.2f\n"+
			"[yellow]Upper Circuit:[white] %.2f",
		data.LastPrice,
		data.OHLC.Open, data.OHLC.High,
		data.OHLC.Low, data.OHLC.Close,
		data.Volume,
		data.OI,
		data.NetChange,
		data.LowerCircuitLimit,
		data.UpperCircuitLimit,
	))

	// ======================================
	// BUY TABLE
	// ======================================
	depthBuyTable := tview.NewTable()
	depthBuyTable.SetTitle(" BUY Depth ")
	depthBuyTable.SetBorder(true)

	headers := []string{"Type", "Price", "Qty", "Orders"}

	// BUY headers
	for i, h := range headers {
		depthBuyTable.SetCell(0, i, tview.NewTableCell("[yellow::b]"+h))
	}

	// BUY rows
	for r := 0; r < 5; r++ {
		b := data.Depth.Buy[r]
		row := r + 1
		depthBuyTable.SetCell(row, 0, cell("[green]BUY"))
		depthBuyTable.SetCell(row, 1, cellFloat(b.Price))
		depthBuyTable.SetCell(row, 2, cellInt(b.Quantity))
		depthBuyTable.SetCell(row, 3, cellInt(b.Orders))
	}

	// ======================================
	// SELL TABLE
	// ======================================
	depthSellTable := tview.NewTable()

	depthSellTable.SetTitle(" SELL Depth ")
	depthSellTable.SetBorder(true)

	// SELL headers
	for i, h := range headers {
		depthSellTable.SetCell(0, i, tview.NewTableCell("[yellow::b]"+h))
	}

	// SELL rows
	for r := 0; r < 5; r++ {
		s := data.Depth.Sell[r]
		row := r + 1
		depthSellTable.SetCell(row, 0, cell("[red]SELL"))
		depthSellTable.SetCell(row, 1, cellFloat(s.Price))
		depthSellTable.SetCell(row, 2, cellInt(s.Quantity))
		depthSellTable.SetCell(row, 3, cellInt(s.Orders))
	}

	// ======================================
	// ALIGN TABLES SIDE BY SIDE (VERTICAL PARALLEL)
	// ======================================
	depthColumns := tview.NewFlex()

	depthColumns.SetDirection(tview.FlexColumn)
	depthColumns.AddItem(depthBuyTable, 0, 1, false)
	depthColumns.AddItem(depthSellTable, 0, 1, false)

	// ======================================
	// FINAL LAYOUT (LTP left | DEPTH right)
	// ======================================
	layout := tview.NewFlex()

	layout.SetDirection(tview.FlexColumn)
	layout.AddItem(ltpPanel, 40, 0, false) // fixed width panel
	layout.AddItem(depthColumns, 0, 1, true)

	return layout
}

func cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).SetAlign(tview.AlignCenter)
}

func cellInt(n int) *tview.TableCell {
	return tview.NewTableCell(fmt.Sprintf("%d", n)).SetAlign(tview.AlignCenter)
}

func cellFloat(f float64) *tview.TableCell {
	return tview.NewTableCell(fmt.Sprintf("%.2f", f)).SetAlign(tview.AlignCenter)
}
