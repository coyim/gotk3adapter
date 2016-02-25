package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type treeView struct {
	*gtk.TreeView
}

func wrapTreeView(v *gtk.TreeView, e error) (*treeView, error) {
	if v == nil {
		return nil, e
	}
	return &treeView{v}, e
}

func unwrapTreeView(v gtki.TreeView) *gtk.TreeView {
	if v == nil {
		return nil
	}
	return v.(*treeView).TreeView
}
