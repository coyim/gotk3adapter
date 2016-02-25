package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type checkButton struct {
	*gtk.CheckButton
}

func wrapCheckButton(v *gtk.CheckButton, e error) (*checkButton, error) {
	if v == nil {
		return nil, e
	}
	return &checkButton{v}, e
}

func unwrapCheckButton(v gtki.CheckButton) *gtk.CheckButton {
	if v == nil {
		return nil
	}
	return v.(*checkButton).CheckButton
}
