package gtka

import "github.com/gotk3/gotk3/gtk"

type aboutDialog struct {
	*gtk.AboutDialog
}

func wrapAboutDialog(v *gtk.AboutDialog, e error) (*aboutDialog, error) {
	if v == nil {
		return nil, e
	}
	return &aboutDialog{v}, e
}
