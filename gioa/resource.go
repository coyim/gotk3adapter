package gioa

import (
	"github.com/coyim/gotk3adapter/gioi"
	"github.com/gotk3/gotk3/gio"
)

type resource struct {
	internal gio.GResource
}

func WrapResourceSimple(v gio.GResource) gioi.Resource {
	if v == nil {
		return nil
	}
	return &resource{v}
}

func WrapResource(v gio.GResource, e error) (gioi.Resource, error) {
	return WrapResourceSimple(v), e
}

func UnwrapResource(v gioi.Resource) gio.GResource {
	if v == nil {
		return nil
	}
	return v.(*resource).internal
}
