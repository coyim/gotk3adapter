package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type builder struct {
	*gtk.Builder
}

func wrapBuilder(v *gtk.Builder, e error) (*builder, error) {
	if v == nil {
		return nil, e
	}
	return &builder{v}, e
}

func unwrapBuilder(v gtki.Builder) *gtk.Builder {
	if v == nil {
		return nil
	}
	return v.(*builder).Builder
}
