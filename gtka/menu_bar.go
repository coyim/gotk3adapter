package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type menuBar struct {
	*gtk.MenuBar
}

func wrapMenuBar(v *gtk.MenuBar, e error) (*menuBar, error) {
	if v == nil {
		return nil, e
	}
	return &menuBar{v}, e
}

func unwrapMenuBar(v gtki.MenuBar) *gtk.MenuBar {
	if v == nil {
		return nil
	}
	return v.(*menuBar).MenuBar
}
