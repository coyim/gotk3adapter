package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type messageDialog struct {
	*gtk.MessageDialog
}

func wrapMessageDialog(v *gtk.MessageDialog, e error) (*messageDialog, error) {
	if v == nil {
		return nil, e
	}
	return &messageDialog{v}, e
}

func unwrapMessageDialog(v gtki.MessageDialog) *gtk.MessageDialog {
	if v == nil {
		return nil
	}
	return v.(*messageDialog).MessageDialog
}
