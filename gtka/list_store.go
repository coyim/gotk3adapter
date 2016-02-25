package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type listStore struct {
	*gtk.ListStore
}

func wrapListStore(v *gtk.ListStore, e error) (*listStore, error) {
	if v == nil {
		return nil, e
	}
	return &listStore{v}, e
}

func unwrapListStore(v gtki.ListStore) *gtk.ListStore {
	if v == nil {
		return nil
	}
	return v.(*listStore).ListStore
}
