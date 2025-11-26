package main

import "KiteRunner/internal/app"

func main() {
	a := app.New()
	if err := a.Run(); err != nil {
		panic(err)
	}
}
