package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			myWindow.Close()
		},
	}

	// we can also append items
	form.Append("Text", textArea)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
