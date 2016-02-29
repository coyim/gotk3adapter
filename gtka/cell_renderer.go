package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gliba"
	"github.com/twstrike/gotk3adapter/gtki"
)

type cellRenderer struct {
	*gliba.Object
	*gtk.CellRenderer
}

func wrapCellRendererSimple(v *gtk.CellRenderer) *cellRenderer {
	if v == nil {
		return nil
	}
	return &cellRenderer{gliba.WrapObjectSimple(v.Object), v}
}

func wrapCellRenderer(v *gtk.CellRenderer, e error) (*cellRenderer, error) {
	return wrapCellRendererSimple(v), e
}

func unwrapCellRenderer(v gtki.CellRenderer) *gtk.CellRenderer {
	if v == nil {
		return nil
	}
	return v.(*cellRenderer).CellRenderer
}
