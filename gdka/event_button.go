package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type eventButton struct {
	*gdk.EventButton
}

func wrapEventButton(v *gdk.EventButton, e error) (*eventButton, error) {
	if v == nil {
		return nil, e
	}
	return &eventButton{v}, e
}

func unwrapEventButton(v gdki.EventButton) *gdk.EventButton {
	if v == nil {
		return nil
	}
	return v.(*eventButton).EventButton
}
