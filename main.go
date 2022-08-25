package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("calc")
	label := widget.NewLabel("calc")

	entryA := widget.NewEntry()
	entryB := widget.NewEntry()

	btn := widget.NewButton("+", func() {
		n1, err1 := strconv.ParseFloat(entryA.Text, 64)
		n2, err2 := strconv.ParseFloat(entryB.Text, 64)
		if err1 != nil || err2 != nil {
			label.SetText("input error")
			return
		}
		sum := n1 + n2
		label.SetText(fmt.Sprintf("%f", sum))
	})

	w.SetContent(container.NewVBox(
		label,
		entryA,
		entryB,
		btn,
	))

	w.ShowAndRun()
}
