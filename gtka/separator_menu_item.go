package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type separatorMenuItem struct {
	*gtk.SeparatorMenuItem
}

func wrapSeparatorMenuItem(v *gtk.SeparatorMenuItem, e error) (*separatorMenuItem, error) {
	if v == nil {
		return nil, e
	}
	return &separatorMenuItem{v}, e
}

func unwrapSeparatorMenuItem(v gtki.SeparatorMenuItem) *gtk.SeparatorMenuItem {
	if v == nil {
		return nil
	}
	return v.(*separatorMenuItem).SeparatorMenuItem
}
