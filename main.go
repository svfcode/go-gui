package main

import (
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	
	w := a.NewWindow("calc")
	w.Resize(fyne.NewSize(400, 300))

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

	checks := widget.NewCheckGroup([]string{"check 1", "check 2"}, func(s []string) {})
	btnForChecks := widget.NewButton("press", func() {
		for _, i := range checks.Selected {
			fmt.Println(i)
		}
	})

	url, _ := url.Parse("ya.ru")
	link := widget.NewHyperlink("ya", url)

	w.SetContent(container.NewVBox(
		label,
		entryA,
		entryB,
		btn,
		checks,
		btnForChecks,
		link,
	))

	w.ShowAndRun()
}
