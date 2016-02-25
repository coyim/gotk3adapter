package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type button struct {
	*gtk.Button
}

func wrapButton(v *gtk.Button, e error) (*button, error) {
	if v == nil {
		return nil, e
	}
	return &button{v}, e
}

func unwrapButton(v gtki.Button) *gtk.Button {
	if v == nil {
		return nil
	}
	return v.(*button).Button
}
