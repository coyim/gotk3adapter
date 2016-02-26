package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/gotk3adapter/gliba"
	"github.com/twstrike/gotk3adapter/gtki"
)

type widget struct {
	*gliba.Object
	*gtk.Widget
}

func wrapWidgetSimple(v *gtk.Widget) *widget {
	if v == nil {
		return nil
	}
	return &widget{gliba.WrapObjectSimple(&v.InitiallyUnbound.Object), v}
}

func wrapWidget(v *gtk.Widget, e error) (*widget, error) {
	return wrapWidgetSimple(v), e
}

func unwrapWidget(v gtki.Widget) *gtk.Widget {
	if v == nil {
		return nil
	}
	return v.(*widget).Widget
}

func (v *widget) GetWindow() (gdki.Window, error) {
	return gdka.WrapWindow(v.Widget.GetWindow())
}

func (v *widget) GetStyleContext() (gtki.StyleContext, error) {
	return wrapStyleContext(v.Widget.GetStyleContext())
}

func (v *widget) SetHAlign(v gtki.Align) {
	v.Widget.SetHAlign(gtk.Align(v))
}
