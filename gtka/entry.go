package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type entry struct {
	*gtk.Entry
}

func wrapEntry(v *gtk.Entry, e error) (*entry, error) {
	if v == nil {
		return nil, e
	}
	return &entry{v}, e
}

func unwrapEntry(v gtki.Entry) *gtk.Entry {
	if v == nil {
		return nil
	}
	return v.(*entry).Entry
}
