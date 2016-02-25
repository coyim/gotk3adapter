package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type comboBox struct {
	*gtk.ComboBox
}

func wrapComboBox(v *gtk.ComboBox, e error) (*comboBox, error) {
	if v == nil {
		return nil, e
	}
	return &comboBox{v}, e
}

func unwrapComboBox(v gtki.ComboBox) *gtk.ComboBox {
	if v == nil {
		return nil
	}
	return v.(*comboBox).ComboBox
}
