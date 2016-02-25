package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type cellRendererToggle struct {
	*gtk.CellRendererToggle
}

func wrapCellRendererToggle(v *gtk.CellRendererToggle, e error) (*cellRendererToggle, error) {
	if v == nil {
		return nil, e
	}
	return &cellRendererToggle{v}, e
}

func unwrapCellRendererToggle(v gtki.CellRendererToggle) *gtk.CellRendererToggle {
	if v == nil {
		return nil
	}
	return v.(*cellRendererToggle).CellRendererToggle
}
