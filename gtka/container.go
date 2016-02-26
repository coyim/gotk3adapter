package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type container struct {
	*widget
	*gtk.Container
}

func wrapContainerSimple(v *gtk.Container) *container {
	if v == nil {
		return nil
	}
	return &container{wrapWidgetSimple(&v.Widget), v}
}

func wrapContainer(v *gtk.Container, e error) (*container, error) {
	return wrapContainerSimple(v), e
}

func unwrapContainer(v gtki.Container) *gtk.Container {
	if v == nil {
		return nil
	}
	return v.(*container).Container
}

// TODO:
// Add(Widget)
// Remove(Widget)
