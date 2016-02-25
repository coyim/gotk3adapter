package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type cssProvider struct {
	*gtk.CssProvider
}

func wrapCssProvider(v *gtk.CssProvider, e error) (*cssProvider, error) {
	if v == nil {
		return nil, e
	}
	return &cssProvider{v}, e
}

func unwrapCssProvider(v gtki.CssProvider) *gtk.CssProvider {
	if v == nil {
		return nil
	}
	return v.(*cssProvider).CssProvider
}
