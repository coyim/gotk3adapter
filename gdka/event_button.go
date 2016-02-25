package gdka

import "github.com/gotk3/gotk3/gdk"

type eventButton struct {
	*gdk.EventButton
}

func wrapEventAsEventButton(v *event) *eventButton {
	return wrapEventButton(&gdk.EventButton{v.Event})
}

func wrapEventButton(v *gdk.EventButton) *eventButton {
	return &eventButton{v}
}
