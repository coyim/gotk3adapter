package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type treeStore struct {
	*gtk.TreeStore
}

func wrapTreeStore(v *gtk.TreeStore, e error) (*treeStore, error) {
	if v == nil {
		return nil, e
	}
	return &treeStore{v}, e
}

func unwrapTreeStore(v gtki.TreeStore) *gtk.TreeStore {
	if v == nil {
		return nil
	}
	return v.(*treeStore).TreeStore
}
