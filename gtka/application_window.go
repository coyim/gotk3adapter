package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type applicationWindow struct {
	*gtk.ApplicationWindow
}

func wrapApplicationWindow(v *gtk.ApplicationWindow, e error) (*applicationWindow, error) {
	if v == nil {
		return nil, e
	}
	return &applicationWindow{v}, e
}

func unwrapApplicationWindow(v gtki.ApplicationWindow) *gtk.ApplicationWindow {
	if v == nil {
		return nil
	}
	return v.(*applicationWindow).ApplicationWindow
}
