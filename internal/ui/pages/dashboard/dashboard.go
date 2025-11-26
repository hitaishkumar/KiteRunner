package pages

import (
	"fmt"

	"KiteRunner/internal/model"
	"KiteRunner/internal/ui/banner"
	"KiteRunner/internal/ui/layout"

	"github.com/rivo/tview"
)

func Dashboard(a *model.App) tview.Primitive {

	userData := a.UserProfile // ← Clean + works

	// -----------------------------
	// Banner
	// -----------------------------
	bannerView := tview.NewTextView()
	bannerView.SetText(banner.SmallBanner())

	// ============================================================
	// BANK ACCOUNTS TABLE
	// ============================================================
	bankTable := tview.NewTable()

	bankTable.SetCell(0, 0, cellHeader("Bank"))
	bankTable.SetCell(0, 1, cellHeader("Account"))

	for i, b := range userData.Data.BankAccounts {
		row := i + 1
		bankTable.SetCell(row, 0, cell(b.Name))
		bankTable.SetCell(row, 1, cell(b.Account))
	}

	// ============================================================
	// EXCHANGES TABLE
	// ============================================================
	exchTable := tview.NewTable()

	exchTable.SetCell(0, 0, cellHeader("Exchange"))
	for i, ex := range userData.Data.Exchanges {
		exchTable.SetCell(i+1, 0, cell(ex))
	}

	// ============================================================
	// PRODUCTS / ORDER TYPES TABLE
	// ============================================================
	prodTable := tview.NewTable()

	prodTable.SetCell(0, 0, cellHeader("Products"))
	prodTable.SetCell(0, 1, cellHeader("Order Types"))

	maxLen := max(len(userData.Data.Products), len(userData.Data.OrderTypes))

	for i := 0; i < maxLen; i++ {
		if i < len(userData.Data.Products) {
			prodTable.SetCell(i+1, 0, cell(userData.Data.Products[i]))
		}
		if i < len(userData.Data.OrderTypes) {
			prodTable.SetCell(i+1, 1, cell(userData.Data.OrderTypes[i]))
		}
	}

	// ============================================================
	// DP IDs List
	// ============================================================
	dpList := tview.NewList()

	for _, dp := range userData.Data.DpIds {
		dpList.AddItem(dp, "", 0, nil)
	}

	// ============================================================
	// META PANEL
	// ============================================================
	meta := tview.NewTextView()

	meta.SetText(
		fmt.Sprintf("[blue]POA: [white] %s\n", userData.Data.Meta.Poa) +
			fmt.Sprintf("[blue]Silo: [white] %s\n", userData.Data.Meta.Silo),
	)

	// ============================================================
	// 2FA PANEL
	// ============================================================
	twofa := tview.NewTextView()
	twofa.SetText(
		fmt.Sprintf("[blue]Type: [white] %s\n", userData.Data.TwofaType) +
			fmt.Sprintf("[blue]Timestamp: [white] %s\n", userData.Data.TwofaTimestamp),
	)

	// ============================================================
	// GRID LAYOUT
	// ============================================================
	grid := tview.NewGrid().
		SetRows(7, 0, 0).
		SetColumns(0, 0, 0).
		SetBorders(true)

	// Row 0: Full width userPanel

	// Row 1: 3 tables
	grid.AddItem(bankTable, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(exchTable, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(prodTable, 1, 2, 1, 1, 0, 0, false)

	// Row 2: misc items
	grid.AddItem(dpList, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(meta, 2, 1, 1, 1, 0, 0, false)
	grid.AddItem(twofa, 2, 2, 1, 1, 0, 0, false)

	return layout.WithContextBanner(bannerView, grid, a)
}

// ============================================================
// Helper functions
// ============================================================

func cellHeader(text string) *tview.TableCell {
	return tview.NewTableCell("[yellow::b]" + text).
		SetAlign(tview.AlignCenter)
}

func cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetAlign(tview.AlignCenter)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
