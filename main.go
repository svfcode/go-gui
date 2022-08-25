package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("first app")
	labal := widget.NewLabel("hiasdf dsfa adsf")

	w.SetContent(container.NewVBox(
		labal,
	))

	w.ShowAndRun()
}