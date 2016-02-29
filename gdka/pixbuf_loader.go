package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/gotk3adapter/gliba"
)

type pixbufLoader struct {
	*gliba.Object
	internal *gdk.PixbufLoader
}

func wrapPixbufLoader(v *gdk.PixbufLoader, e error) (*pixbufLoader, error) {
	if v == nil {
		return nil, e
	}
	return &pixbufLoader{gliba.WrapObjectSimple(v.Object), v}, e
}

func unwrapPixbufLoader(v gdki.PixbufLoader) *gdk.PixbufLoader {
	if v == nil {
		return nil
	}
	return v.(*pixbufLoader).internal
}

func (v *pixbufLoader) Close() error {
	return v.internal.Close()
}

func (v *pixbufLoader) GetPixbuf() (gdki.Pixbuf, error) {
	return wrapPixbuf(v.internal.GetPixbuf())
}

func (v *pixbufLoader) Write(b []byte) (int, error) {
	return v.internal.Write(b)
}
