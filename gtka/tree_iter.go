package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type treeIter struct {
	*gtk.TreeIter
}

func wrapTreeIter(v *gtk.TreeIter, e error) (*treeIter, error) {
	if v == nil {
		return nil, e
	}
	return &treeIter{v}, e
}

func unwrapTreeIter(v gtki.TreeIter) *gtk.TreeIter {
	if v == nil {
		return nil
	}
	return v.(*treeIter).TreeIter
}
