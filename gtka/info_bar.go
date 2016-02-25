package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type infoBar struct {
	*gtk.InfoBar
}

func wrapInfoBar(v *gtk.InfoBar, e error) (*infoBar, error) {
	if v == nil {
		return nil, e
	}
	return &infoBar{v}, e
}

func unwrapInfoBar(v gtki.InfoBar) *gtk.InfoBar {
	if v == nil {
		return nil
	}
	return v.(*infoBar).InfoBar
}
