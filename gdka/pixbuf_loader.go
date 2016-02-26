package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/gotk3adapter/gliba"
)

type pixbufLoader struct {
	*gliba.Object
	*gdk.PixbufLoader
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
	return v.(*pixbufLoader).PixbufLoader
}

// Specific implementations

// func (v *pixbufLoader) Close() error {
// 	return v.Close()
// }

func (v *pixbufLoader) GetPixbuf() (gdki.Pixbuf, error) {
	return wrapPixbuf(v.PixbufLoader.GetPixbuf())
}

// func (v *pixbufLoader) Write(b []byte) (int, error) {
// 	return v.Write(b)
// }

// Object implementations

// func (v *pixbufLoader) Connect(s string, v1 interface{}, v2 ...interface{}) (glibi.SignalHandle, error) {
// 	return gliba.Object_Connect(v.PixbufLoader, s, v1, v2...)
// }

// func (v *pixbufLoader) Emit(s string, v1 ...interface{}) (interface{}, error) {
// 	return gliba.Object_Emit(v.PixbufLoader, s, v1...)
// }

// func (v *pixbufLoader) SetProperty(s string, v1 interface{}) error {
// 	return gliba.Object_SetProperty(v.PixbufLoader, s, v1)
// }
