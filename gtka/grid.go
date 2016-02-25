package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type grid struct {
	*gtk.Grid
}

func wrapGrid(v *gtk.Grid, e error) (*grid, error) {
	if v == nil {
		return nil, e
	}
	return &grid{v}, e
}

func unwrapGrid(v gtki.Grid) *gtk.Grid {
	if v == nil {
		return nil
	}
	return v.(*grid).Grid
}
