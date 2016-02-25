package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type widget struct {
	*gtk.Widget
}

func wrapWidget(v *gtk.Widget, e error) (*widget, error) {
	if v == nil {
		return nil, e
	}
	return &widget{v}, e
}

func unwrapWidget(v gtki.Widget) *gtk.Widget {
	if v == nil {
		return nil
	}
	return v.(*widget).Widget
}
