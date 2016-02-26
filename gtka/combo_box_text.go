package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type comboBoxText struct {
	*comboBox
	*gtk.ComboBoxText
}

func wrapComboBoxTextSimple(v *gtk.ComboBoxText) *comboBoxText {
	if v == nil {
		return nil
	}
	return &comboBoxText{wrapComboBoxSimple(&v.ComboBox), v}
}

func wrapComboBoxText(v *gtk.ComboBoxText, e error) (*comboBoxText, error) {
	return wrapComboBoxTextSimple(v), e
}

func unwrapComboBoxText(v gtki.ComboBoxText) *gtk.ComboBoxText {
	if v == nil {
		return nil
	}
	return v.(*comboBoxText).ComboBoxText
}
