type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}

type Button struct {
	BaseWidget
	Text  string
	Style ButtonStyle
	Icon  fyne.Resource

	OnTapped func()
}

type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size

	Refresh()
	Objects() []CanvasObject
	Destroy()
}

type buttonRenderer struct {
	icon   *canvas.Image
	label  *canvas.Text
	shadow *fyne.CanvasObject

	objects []fyne.CanvasObject
	button  *Button
}