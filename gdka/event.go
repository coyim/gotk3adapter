package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type event struct {
	*gdk.Event
}

func eventCast(e gdki.Event) *event {
	if e == nil {
		return nil
	}
	return e.(*event)
}
