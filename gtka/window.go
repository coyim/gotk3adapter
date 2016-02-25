package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type window struct {
	*gtk.Window
}

func wrapWindow(v *gtk.Window, e error) (*window, error) {
	if v == nil {
		return nil, e
	}
	return &window{v}, e
}

func unwrapWindow(v gtki.Window) *gtk.Window {
	if v == nil {
		return nil
	}
	return v.(*window).Window
}
