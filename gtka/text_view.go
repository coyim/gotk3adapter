package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type textView struct {
	*gtk.TextView
}

func wrapTextView(v *gtk.TextView, e error) (*textView, error) {
	if v == nil {
		return nil, e
	}
	return &textView{v}, e
}

func unwrapTextView(v gtki.TextView) *gtk.TextView {
	if v == nil {
		return nil
	}
	return v.(*textView).TextView
}
