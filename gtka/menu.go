package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type menu struct {
	*gtk.Menu
}

func wrapMenu(v *gtk.Menu, e error) (*menu, error) {
	if v == nil {
		return nil, e
	}
	return &menu{v}, e
}

func unwrapMenu(v gtki.Menu) *gtk.Menu {
	if v == nil {
		return nil
	}
	return v.(*menu).Menu
}
