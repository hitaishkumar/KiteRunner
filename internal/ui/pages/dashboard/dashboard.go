package pages

import (
	"fmt"

	"KiteRunner/internal/model"
	"KiteRunner/internal/ui/banner"
	"KiteRunner/internal/ui/layout"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Dashboard(a *model.App) tview.Primitive {
	user := a.UserProfile.Data

	bannerView := tview.NewTextView().
		SetText(banner.SmallBanner()).
		SetDynamicColors(true)

	// ---------------------------------------------------------
	// COMPACT TABLES
	// ---------------------------------------------------------
	bankTable := compactTable(
		[]string{"Bank", "Account"},
		func() int { return len(user.BankAccounts) },
		func(i int) []string {
			b := user.BankAccounts[i]
			return []string{
				trim(b.Name, 12),
				trim(b.Account, 10),
			}
		})

	exchTable := compactTable(
		[]string{"Exchange"},
		func() int { return len(user.Exchanges) },
		func(i int) []string {
			return []string{trim(user.Exchanges[i], 12)}
		})

	prodTable := compactTwoColumn(
		"Prod", "Order",
		user.Products,
		user.OrderTypes,
	)

	// ---------------------------------------------------------
	// COMPACT LISTS
	// ---------------------------------------------------------
	dpList := tview.NewList().ShowSecondaryText(false)
	for _, dp := range user.DpIds {
		dpList.AddItem(trim(dp, 14), "", 0, nil)
	}

	meta := compactInfo(map[string]string{
		"POA":  user.Meta.Poa,
		"Silo": user.Meta.Silo,
	})

	twofa := compactInfo(map[string]string{
		"Type": user.TwofaType,
		"Time": user.TwofaTimestamp,
	})

	// ---------------------------------------------------------
	// GRID (unchanged)
	// ---------------------------------------------------------
	grid := tview.NewGrid().
		SetRows(3).
		SetColumns(0, 0)

	grid.SetBorders(true).SetBorderColor(tcell.ColorDarkRed)

	grid.AddItem(bankTable, 1, 0, 1, 1, 0, 0, true)
	grid.AddItem(exchTable, 1, 1, 1, 1, 0, 0, true)
	grid.AddItem(prodTable, 1, 2, 1, 1, 0, 0, true)

	grid.AddItem(dpList, 2, 0, 1, 1, 0, 0, true)
	grid.AddItem(meta, 2, 1, 1, 1, 0, 0, true)
	grid.AddItem(twofa, 2, 2, 1, 1, 0, 0, true)

	return layout.WithContextBanner(bannerView, grid, a)
}

//
// ----------------------------------------------------------
// COMPACT HELPERS
// ----------------------------------------------------------
//

func compactTable(headers []string, rows func() int, data func(i int) []string) *tview.Table {
	t := tview.NewTable().
		SetBorders(true)

	// header
	for c, h := range headers {
		t.SetCell(0, c,
			tview.NewTableCell("[yellow::b]"+h).
				SetAlign(tview.AlignLeft).
				SetExpansion(2),
		)
	}

	// rows
	for i := 0; i < rows(); i++ {
		cols := data(i)
		for c, v := range cols {
			t.SetCell(i+1, c,
				tview.NewTableCell(v).
					SetAlign(tview.AlignLeft))
		}
	}

	return t
}

func compactTwoColumn(h1, h2 string, left, right []string) *tview.Table {
	t := tview.NewTable()

	t.SetCell(0, 0, compactHeader(h1))
	t.SetCell(0, 1, compactHeader(h2))

	n := max(len(left), len(right))
	for i := 0; i < n; i++ {
		if i < len(left) {
			t.SetCell(i+1, 0, compactCell(trim(left[i], 12)))
		}
		if i < len(right) {
			t.SetCell(i+1, 1, compactCell(trim(right[i], 12)))
		}
	}

	return t
}

func compactHeader(text string) *tview.TableCell {
	return tview.NewTableCell("[yellow::b]" + text).
		SetAlign(tview.AlignLeft)
}

func compactCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetAlign(tview.AlignLeft)
}

func compactInfo(m map[string]string) *tview.TextView {
	tv := tview.NewTextView().
		SetDynamicColors(true)

	for k, v := range m {
		tv.Write([]byte(fmt.Sprintf("[blue]%s:[white] %s\n", k, trim(v, 14))))
	}

	return tv
}

func trim(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
