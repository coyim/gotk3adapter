package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type pixbuf struct {
	*gdk.Pixbuf
}

func wrapPixbuf(v *gdk.Pixbuf, e error) (*pixbuf, error) {
	if v == nil {
		return nil, e
	}
	return &pixbuf{v}, e
}

func UnwrapPixbuf(v gdki.Pixbuf) *gdk.Pixbuf {
	if v == nil {
		return nil
	}
	return v.(*pixbuf).Pixbuf
}
