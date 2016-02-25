package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type checkMenuItem struct {
	*gtk.CheckMenuItem
}

func wrapCheckMenuItem(v *gtk.CheckMenuItem, e error) (*checkMenuItem, error) {
	if v == nil {
		return nil, e
	}
	return &checkMenuItem{v}, e
}

func unwrapCheckMenuItem(v gtki.CheckMenuItem) *gtk.CheckMenuItem {
	if v == nil {
		return nil
	}
	return v.(*checkMenuItem).CheckMenuItem
}
