package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type builder struct {
	*gliba.Object
	*gtk.Builder
}

func wrapBuilderSimple(v *gtk.Builder) *builder {
	if v == nil {
		return nil
	}
	return &builder{gliba.WrapObjectSimple(&v.Object), v}
}

func wrapBuilder(v *gtk.Builder, e error) (*builder, error) {
	return wrapBuilderSimple(v), e
}

func unwrapBuilder(v gtki.Builder) *gtk.Builder {
	if v == nil {
		return nil
	}
	return v.(*builder).Builder
}
