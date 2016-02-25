package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type pixbufLoader struct {
	*gdk.PixbufLoader
}

func wrapPixbufLoader(v *gdk.PixbufLoader, e error) (*pixbufLoader, error) {
	if v == nil {
		return nil, e
	}
	return &pixbufLoader{v}, e
}

func unwrapPixbufLoader(v gdki.PixbufLoader) *gdk.PixbufLoader {
	if v == nil {
		return nil
	}
	return v.(*pixbufLoader).PixbufLoader
}
