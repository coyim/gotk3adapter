package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type headerBar struct {
	*gtk.HeaderBar
}

func wrapHeaderBar(v *gtk.HeaderBar, e error) (*headerBar, error) {
	if v == nil {
		return nil, e
	}
	return &headerBar{v}, e
}

func unwrapHeaderBar(v gtki.HeaderBar) *gtk.HeaderBar {
	if v == nil {
		return nil
	}
	return v.(*headerBar).HeaderBar
}
