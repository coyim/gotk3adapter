package gioa

import (
	"github.com/coyim/gotk3adapter/gioi"
	"github.com/gotk3/gotk3/gio"
)

type RealGio struct{}

var Real = &RealGio{}

func (*RealGio) LoadResource(path string) (gioi.Resource, error) {
	return WrapResource(gio.LoadGResource(path))
}

func (*RealGio) NewResourceFromData(data []byte) (gioi.Resource, error) {
	return WrapResource(gio.NewGResourceFromData(data))
}

func (*RealGio) RegisterResource(r gioi.Resource) {
	gio.RegisterGResource(UnwrapResource(r))
}

func (*RealGio) UnregisterResource(r gioi.Resource) {
	gio.UnregisterGResource(UnwrapResource(r))
}
