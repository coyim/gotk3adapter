package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type comboBox struct {
	*bin
	*gtk.ComboBox
}

func wrapComboBoxSimple(v *gtk.ComboBox) *comboBox {
	if v == nil {
		return nil
	}
	return &comboBox{wrapBin(&v.Bin), v}
}

func wrapComboBox(v *gtk.ComboBox, e error) (*comboBox, error) {
	return wrapComboBoxSimple(v), e
}

func unwrapComboBox(v gtki.ComboBox) *gtk.ComboBox {
	if v == nil {
		return nil
	}
	return v.(*comboBox).ComboBox
}
