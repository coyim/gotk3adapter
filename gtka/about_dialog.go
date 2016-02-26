package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type aboutDialog struct {
	*dialog
	*gtk.AboutDialog
}

func wrapAboutDialog(v *gtk.AboutDialog, e error) (*aboutDialog, error) {
	if v == nil {
		return nil, e
	}
	return &aboutDialog{wrapDialogSimple(&v.Dialog), v}, e
}

func unwrapAboutDialog(v gtki.AboutDialog) *gtk.AboutDialog {
	if v == nil {
		return nil
	}
	return v.(*aboutDialog).AboutDialog
}
