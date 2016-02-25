package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type box struct {
	*gtk.Box
}

func wrapBox(v *gtk.Box, e error) (*box, error) {
	if v == nil {
		return nil, e
	}
	return &box{v}, e
}

func unwrapBox(v gtki.Box) *gtk.Box {
	if v == nil {
		return nil
	}
	return v.(*box).Box
}
