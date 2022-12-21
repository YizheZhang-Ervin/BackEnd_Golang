package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))

	a.Run()
}
