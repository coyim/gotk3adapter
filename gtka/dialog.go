package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type dialog struct {
	*window
	*gtk.Dialog
}

func wrapDialogSimple(v *gtk.Dialog) *dialog {
	if v == nil {
		return nil
	}
	return &dialog{wrapWindowSimple(&v.Window), v}
}

func wrapDialog(v *gtk.Dialog, e error) (*dialog, error) {
	return wrapDialogSimple(v), e
}

func unwrapDialog(v gtki.Dialog) *gtk.Dialog {
	if v == nil {
		return nil
	}
	return v.(*dialog).Dialog
}

// TODO:
// Run() int
// SetDefaultResponse(ResponseType)
