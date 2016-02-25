package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type menuItem struct {
	*gtk.MenuItem
}

func wrapMenuItem(v *gtk.MenuItem, e error) (*menuItem, error) {
	if v == nil {
		return nil, e
	}
	return &menuItem{v}, e
}

func unwrapMenuItem(v gtki.MenuItem) *gtk.MenuItem {
	if v == nil {
		return nil
	}
	return v.(*menuItem).MenuItem
}
