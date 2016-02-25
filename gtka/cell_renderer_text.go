package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type cellRendererText struct {
	*gtk.CellRendererText
}

func wrapCellRendererText(v *gtk.CellRendererText, e error) (*cellRendererText, error) {
	if v == nil {
		return nil, e
	}
	return &cellRendererText{v}, e
}

func unwrapCellRendererText(v gtki.CellRendererText) *gtk.CellRendererText {
	if v == nil {
		return nil
	}
	return v.(*cellRendererText).CellRendererText
}
