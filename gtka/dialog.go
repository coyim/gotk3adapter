package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type dialog struct {
	*gtk.Dialog
}

func wrapDialog(v *gtk.Dialog, e error) (*dialog, error) {
	if v == nil {
		return nil, e
	}
	return &dialog{v}, e
}

func unwrapDialog(v gtki.Dialog) *gtk.Dialog {
	if v == nil {
		return nil
	}
	return v.(*dialog).Dialog
}
