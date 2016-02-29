package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gliba"
)

func init() {
	gliba.AddWrapper(WrapLocal)

	gliba.AddUnwrapper(UnwrapLocal)
}

func WrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *gdk.EventButton:
		return wrapEventButtonSimple(oo), true
	case *gdk.Event:
		return wrapEventSimple(oo), true
	case *gdk.Pixbuf:
		return wrapPixbufSimple(oo), true
	case *gdk.PixbufLoader:
		return wrapPixbufLoaderSimple(oo), true
	case *gdk.Screen:
		return wrapScreenSimple(oo), true
	case *gdk.Window:
		return WrapWindowSimple(oo), true
	default:
		return nil, false
	}
}

func UnwrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *eventButton:
		return unwrapEventButton(oo), true
	case *event:
		return unwrapEvent(oo), true
	case *pixbuf:
		return UnwrapPixbuf(oo), true
	case *pixbufLoader:
		return unwrapPixbufLoader(oo), true
	case *screen:
		return UnwrapScreen(oo), true
	case *window:
		return UnwrapWindow(oo), true
	default:
		return nil, false
	}
}
