package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type textTagTable struct {
	*gtk.TextTagTable
}

func wrapTextTagTable(v *gtk.TextTagTable, e error) (*textTagTable, error) {
	if v == nil {
		return nil, e
	}
	return &textTagTable{v}, e
}

func unwrapTextTagTable(v gtki.TextTagTable) *gtk.TextTagTable {
	if v == nil {
		return nil
	}
	return v.(*textTagTable).TextTagTable
}
