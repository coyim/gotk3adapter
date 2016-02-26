package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type entry struct {
	*widget
	*gtk.Entry
}

func wrapEntrySimple(v *gtk.Entry) *entry {
	if v == nil {
		return nil
	}
	return &entry{wrapWidgetSimple(&v.Widget), v}
}

func wrapEntry(v *gtk.Entry, e error) (*entry, error) {
	return wrapEntrySimple(v), e
}

func unwrapEntry(v gtki.Entry) *gtk.Entry {
	if v == nil {
		return nil
	}
	return v.(*entry).Entry
}
