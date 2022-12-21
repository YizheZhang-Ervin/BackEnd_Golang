package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("com.example.tutorial.preferences")
	w := a.NewWindow("Timeout")

	var timeout time.Duration

	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		a.Preferences().SetString("AppTimeout", selected)
	})

	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

	go func() {
		time.Sleep(timeout)
		a.Quit()
	}()

	w.SetContent(timeoutSelector)
	w.ShowAndRun()
}
