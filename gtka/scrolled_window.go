package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type scrolledWindow struct {
	*gtk.ScrolledWindow
}

func wrapScrolledWindow(v *gtk.ScrolledWindow, e error) (*scrolledWindow, error) {
	if v == nil {
		return nil, e
	}
	return &scrolledWindow{v}, e
}

func unwrapScrolledWindow(v gtki.ScrolledWindow) *gtk.ScrolledWindow {
	if v == nil {
		return nil
	}
	return v.(*scrolledWindow).ScrolledWindow
}
