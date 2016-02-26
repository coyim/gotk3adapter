package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type box struct {
	*container
	*gtk.Box
}

func wrapBoxSimple(v *gtk.Box) *box {
	if v == nil {
		return nil
	}
	return &box{wrapContainerSimple(&v.Container), v}
}

func wrapBox(v *gtk.Box, e error) (*box, error) {
	return wrapBoxSimple(v), e
}

func unwrapBox(v gtki.Box) *gtk.Box {
	if v == nil {
		return nil
	}
	return v.(*box).Box
}
