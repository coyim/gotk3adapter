package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type comboBoxText struct {
	*gtk.ComboBoxText
}

func wrapComboBoxText(v *gtk.ComboBoxText, e error) (*comboBoxText, error) {
	if v == nil {
		return nil, e
	}
	return &comboBoxText{v}, e
}

func unwrapComboBoxText(v gtki.ComboBoxText) *gtk.ComboBoxText {
	if v == nil {
		return nil
	}
	return v.(*comboBoxText).ComboBoxText
}
