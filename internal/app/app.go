package app

import (
	"KiteRunner/internal/config"
	"KiteRunner/internal/model"
	"KiteRunner/internal/router"

	"github.com/rivo/tview"
)

func New() *model.App {

	// Load key mapping config
	if err := config.LoadConfig(); err != nil {
		panic("Cannot load config/keys.json: " + err.Error())
	}

	app := &model.App{
		TUI:         tview.NewApplication(),
		Pages:       tview.NewPages(),
		Mode:        model.ModeNavigation, // <--- ADD
		FooterLeft:  tview.NewTextView(),
		FooterRight: tview.NewTextView(),
	}

	// app.UpdateFooter() <<-- cause go routine deadlock

	// Load user profile JSON
	userData, err := loadFullProfile()
	if err != nil {
		panic("Failed to load full_profile.json: " + err.Error())
	}
	app.UserProfile = userData // store into model.App

	router.Register(app)
	return app
}
