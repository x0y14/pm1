package main

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/x0y14/pm1"
	"os"
)

func main() {
	app := gtk.NewApplication("dev.x0y14.pm1", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		pm1.Activate(app)
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
