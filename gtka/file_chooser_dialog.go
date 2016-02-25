package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type fileChooserDialog struct {
	*gtk.FileChooserDialog
}

func wrapFileChooserDialog(v *gtk.FileChooserDialog, e error) (*fileChooserDialog, error) {
	if v == nil {
		return nil, e
	}
	return &fileChooserDialog{v}, e
}

func unwrapFileChooserDialog(v gtki.FileChooserDialog) *gtk.FileChooserDialog {
	if v == nil {
		return nil
	}
	return v.(*fileChooserDialog).FileChooserDialog
}
