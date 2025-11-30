package router

import (
	"KiteRunner/internal/model"
	dashboard "KiteRunner/internal/ui/pages/dashboard"
	"KiteRunner/internal/ui/pages/instruments"
	login "KiteRunner/internal/ui/pages/login"
	"KiteRunner/internal/ui/pages/orders"
	quotes "KiteRunner/internal/ui/pages/quotes"
	streaming "KiteRunner/internal/ui/pages/streaming"
)

func Register(a *model.App) {
	dashboardPage := dashboard.Dashboard(a)
	loginPage := login.Login(a)
	instrumentPage := instruments.Instruments(a)
	quotesPage := quotes.Quotes(a)

	a.Pages.AddPage("login", loginPage, true, true)
	a.Pages.AddPage("dashboard", dashboardPage, true, false)
	a.Pages.AddPage("instruments", instrumentPage, true, false)
	a.Pages.AddPage("quotes", quotesPage, true, false)
	a.Pages.AddPage("orders", orders.Page(a), true, false)
	a.Pages.AddPage("order_history", orders.Page(a), true, false)
	a.Pages.AddPage("streams", streaming.StreamingChartsPage(a), true, false)

}
