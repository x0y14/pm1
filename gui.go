package pm1

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func Activate(app *gtk.Application) {
	window := gtk.NewApplicationWindow(app)
	window.SetTitle("pm1")
	window.SetChild(gtk.NewLabel("Hello from Go!"))
	window.SetDefaultSize(400, 300)
	window.Show()
}
