package gliba

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/twstrike/gotk3adapter/glibi"
)

type Object struct {
	*glib.Object
}

func WrapObjectSimple(v *glib.Object) *Object {
	if v == nil {
		return nil
	}
	return &Object{v}
}

func (v *Object) Connect(string, interface{}, ...interface{}) (glibi.SignalHandle, error) {
	// TODO: be very careful with this one
	return glibi.SignalHandle(0), nil
}

func (v *Object) Emit(string, ...interface{}) (interface{}, error) {
	return nil, nil
}

func (v *Object) SetProperty(string, interface{}) error {
	// TODO: be very careful with this one
	return nil
}
