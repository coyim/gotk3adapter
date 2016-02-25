package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type revealer struct {
	*gtk.Revealer
}

func wrapRevealer(v *gtk.Revealer, e error) (*revealer, error) {
	if v == nil {
		return nil, e
	}
	return &revealer{v}, e
}

func unwrapRevealer(v gtki.Revealer) *gtk.Revealer {
	if v == nil {
		return nil
	}
	return v.(*revealer).Revealer
}
