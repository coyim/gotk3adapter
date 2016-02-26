package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type checkMenuItem struct {
	*menuItem
	*gtk.CheckMenuItem
}

func wrapCheckMenuItemSimple(v *gtk.CheckMenuItem) *checkMenuItem {
	if v == nil {
		return nil
	}
	return &checkMenuItem{wrapMenuItem(&v.MenuItem), v}
}

func wrapCheckMenuItem(v *gtk.CheckMenuItem, e error) (*checkMenuItem, error) {
	return wrapCheckMenuItemSimple(v), e
}

func unwrapCheckMenuItem(v gtki.CheckMenuItem) *gtk.CheckMenuItem {
	if v == nil {
		return nil
	}
	return v.(*checkMenuItem).CheckMenuItem
}
