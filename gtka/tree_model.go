package gtka

import "github.com/twstrike/gotk3adapter/gtki"
import "github.com/gotk3/gotk3/gtk"

func unwrapTreeModel(s gtki.TreeModel) gtk.ITreeModel {
	if s == nil {
		return nil
	}
	return s.(gtk.ITreeModel)
}

func wrapTreeModelSimple(s gtk.ITreeModel) gtki.TreeModel {
	if s == nil {
		return nil
	}

	switch ss := s.(type) {
	case *gtk.ListStore:
		return wrapListStoreSimple(ss)
	case *gtk.TreeStore:
		return wrapTreeStoreSimple(ss)
	}
	return nil
}
