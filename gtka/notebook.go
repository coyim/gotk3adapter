package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type notebook struct {
	*gtk.Notebook
}

func wrapNotebook(v *gtk.Notebook, e error) (*notebook, error) {
	if v == nil {
		return nil, e
	}
	return &notebook{v}, e
}

func unwrapNotebook(v gtki.Notebook) *gtk.Notebook {
	if v == nil {
		return nil
	}
	return v.(*notebook).Notebook
}
