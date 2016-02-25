package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type event struct {
	*gdk.Event
}

func wrapEvent(v *gdk.Event, e error) (*event, error) {
	if v == nil {
		return nil, e
	}
	return &event{v}, e
}

func unwrapEvent(v gdki.Event) *gdk.Event {
	if v == nil {
		return nil
	}
	return v.(*event).Event
}
