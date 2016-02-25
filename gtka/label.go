package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type label struct {
	*gtk.Label
}

func wrapLabel(v *gtk.Label, e error) (*label, error) {
	if v == nil {
		return nil, e
	}
	return &label{v}, e
}

func unwrapLabel(v gtki.Label) *gtk.Label {
	if v == nil {
		return nil
	}
	return v.(*label).Label
}
