package streaming

import (
	"KiteRunner/internal/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/NimbleMarkets/ntcharts/canvas/runes"
	"github.com/NimbleMarkets/ntcharts/linechart/timeserieslinechart"
	"github.com/rivo/tview"
)

// ---- TVIEW WRAPPER PAGE ----
func StreamingChartsPage(a *model.App) tview.Primitive {
	// ======================================
	// LTP PANEL
	// ======================================
	ltpPanel := tview.NewTextView()
	ltpPanel.SetDynamicColors(true)
	ltpPanel.SetBorder(true)
	ltpPanel.SetTitle(" Quote: " + "symbol")
	ltpPanel.SetText(fmt.Sprintf(
		"[yellow]Last Price:[white] %.2f\n"+
			"[green]Open:[white] %.2f  [green]High:[white] %.2f\n"+
			"[red]Low:[white] %.2f   [red]Close:[white] %.2f\n\n"+
			"[yellow]Volume:[white] %d\n"+
			"[yellow]OI:[white] %d\n"+
			"[yellow]Net Change:[white] %.2f\n\n"+
			"[yellow]Lower Circuit:[white] %.2f\n"+
			"[yellow]Upper Circuit:[white] %.2f",
		"data.LastPrice",
		"data.OHLC", "data.OHLC",
		"data.OHLC", "data.OHLC",
		"data.Volume",
		"data.OI",
		"data.NetChange",
		"data.LowerCircuitLimit",
		"data.UpperCircuitLimit",
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
		row := r + 1
		depthBuyTable.SetCell(row, 0, cell("[green]BUY"))
		depthBuyTable.SetCell(row, 1, cell("b.Price"))
		depthBuyTable.SetCell(row, 2, cell("b.Quantity"))
		depthBuyTable.SetCell(row, 3, cell("b.Orders"))
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
		row := r + 1
		depthSellTable.SetCell(row, 0, cell("[red]SELL"))
		depthSellTable.SetCell(row, 1, cell("s.Price"))
		depthSellTable.SetCell(row, 2, cell("s.Quantity"))
		depthSellTable.SetCell(row, 3, cell("s.Orders"))
	}

	// ======================================
	// ALIGN TABLES SIDE BY SIDE (VERTICAL PARALLEL)
	// ======================================
	depthColumns := tview.NewFlex()
	depthColumns.SetDirection(tview.FlexColumn)
	depthColumns.AddItem(depthBuyTable, 0, 1, false)
	depthColumns.AddItem(depthSellTable, 0, 1, false)

	// ======================================
	// TOP LAYOUT (LTP left | DEPTH right)
	// ======================================
	topLayout := tview.NewFlex()
	topLayout.SetDirection(tview.FlexColumn)
	topLayout.AddItem(ltpPanel, 40, 0, false)
	topLayout.AddItem(depthColumns, 0, 1, false)

	// ======================================
	// TIMELINE CHARTS (3 columns)
	// ======================================

	// Chart 1: Price Movement
	priceChart := createTimeSeriesChart(45, 19, generatePriceData(), runes.ArcLineStyle)
	priceView := tview.NewTextView().SetText(priceChart.View()).SetDynamicColors(true)
	priceView.SetBorder(true).SetTitle(" [blue]Price Movement[white] │ [green]₹ (Braille Chart)")

	// Chart 2: Volume
	volumeChart := createTimeSeriesChart(45, 19, generateVolumeData(), runes.ArcLineStyle)
	volumeView := tview.NewTextView().SetText(volumeChart.View()).SetDynamicColors(true)
	volumeView.SetBorder(true).SetTitle(" [blue]Volume[white] │ [yellow]Units (Braille Chart)")

	// Chart 3: Open Interest
	oiChart := createTimeSeriesChart(45, 19, generateOIData(), runes.ThinLineStyle)
	oiView := tview.NewTextView().SetText(oiChart.View()).SetDynamicColors(true)
	oiView.SetBorder(true).SetTitle(" [blue]Open Interest[white] │ [magenta]Contracts (Braille Chart)")

	// Arrange charts in 3 columns
	chartsRow := tview.NewFlex()
	chartsRow.SetDirection(tview.FlexColumn)
	chartsRow.AddItem(priceView, 0, 1, false)
	chartsRow.AddItem(volumeView, 0, 1, false)
	chartsRow.AddItem(oiView, 0, 1, false)

	// ======================================
	// FINAL LAYOUT (Top section + Charts)
	// ======================================
	mainLayout := tview.NewFlex()
	mainLayout.SetDirection(tview.FlexRow)
	mainLayout.AddItem(topLayout, 0, 1, false)
	mainLayout.AddItem(chartsRow, 0, 1, true)

	return mainLayout
}

// ======================================
// HELPER FUNCTIONS
// ======================================

func cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).SetAlign(tview.AlignCenter)
}

func cellInt(n int) *tview.TableCell {
	return tview.NewTableCell(fmt.Sprintf("%d", n)).SetAlign(tview.AlignCenter)
}

func cellFloat(f float64) *tview.TableCell {
	return tview.NewTableCell(fmt.Sprintf("%.2f", f)).SetAlign(tview.AlignCenter)
}

// ======================================
// CHART CREATION
// ======================================

func createTimeSeriesChart(width, height int, data []timeserieslinechart.TimePoint, _ runes.LineStyle) timeserieslinechart.Model {
	// Create chart with default settings
	// Note: timeserieslinechart only supports DrawBraille() method
	chart := timeserieslinechart.New(width, height)

	// Push all data points
	for _, point := range data {
		chart.Push(point)
	}

	// Draw using braille (only available method for timeserieslinechart)
	chart.DrawBraille()

	return chart
}

// ======================================
// DATA GENERATION FUNCTIONS
// ======================================

func generatePriceData() []timeserieslinechart.TimePoint {
	rand.Seed(time.Now().UnixNano())
	data := make([]timeserieslinechart.TimePoint, 50)
	basePrice := 1000.0
	currentTime := time.Now().Add(-50 * time.Minute)

	for i := 0; i < 50; i++ {
		change := (rand.Float64() - 0.5) * 10 // Random change between -5 and +5
		basePrice += change
		data[i] = timeserieslinechart.TimePoint{
			Time:  currentTime.Add(time.Duration(i) * time.Minute),
			Value: basePrice,
		}
	}
	return data
}

func generateVolumeData() []timeserieslinechart.TimePoint {
	rand.Seed(time.Now().UnixNano() + 1)
	data := make([]timeserieslinechart.TimePoint, 50)
	currentTime := time.Now().Add(-50 * time.Minute)

	for i := 0; i < 50; i++ {
		volume := float64(rand.Intn(100000) + 50000) // Random volume between 50k and 150k
		data[i] = timeserieslinechart.TimePoint{
			Time:  currentTime.Add(time.Duration(i) * time.Minute),
			Value: volume,
		}
	}
	return data
}

func generateOIData() []timeserieslinechart.TimePoint {
	rand.Seed(time.Now().UnixNano() + 2)
	data := make([]timeserieslinechart.TimePoint, 50)
	baseOI := 500000.0
	currentTime := time.Now().Add(-50 * time.Minute)

	for i := 0; i < 50; i++ {
		change := (rand.Float64() - 0.5) * 5000 // Gradual OI changes
		baseOI += change
		data[i] = timeserieslinechart.TimePoint{
			Time:  currentTime.Add(time.Duration(i) * time.Minute),
			Value: baseOI,
		}
	}
	return data
}
