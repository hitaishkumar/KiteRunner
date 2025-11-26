package pages

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// -------------------------------
// Helper to fill tables
// -------------------------------
func fillRandomTable(t *tview.Table, title string) {

	// Header cell
	t.SetCell(0, 0,
		tview.NewTableCell(title).
			SetSelectable(false).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter),
	)

	// Random rows
	for r := 1; r <= 5; r++ {
		for c := 0; c < 3; c++ {
			val := rand.Intn(900) + 100 // random 100–999
			t.SetCell(r, c,
				tview.NewTableCell(
					formatNumber(val),
				).SetAlign(tview.AlignCenter),
			)
		}
	}
}

func formatNumber(n int) string {
	return tview.Escape(
		fmt.Sprintf("%03d", n),
	)
}
