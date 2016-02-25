package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type textTag struct {
	*gtk.TextTag
}

func wrapTextTag(v *gtk.TextTag, e error) (*textTag, error) {
	if v == nil {
		return nil, e
	}
	return &textTag{v}, e
}

func unwrapTextTag(v gtki.TextTag) *gtk.TextTag {
	if v == nil {
		return nil
	}
	return v.(*textTag).TextTag
}
