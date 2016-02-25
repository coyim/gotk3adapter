package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type accelGroup struct {
	*gtk.AccelGroup
}

func wrapAccelGroup(v *gtk.AccelGroup, e error) (*accelGroup, error) {
	if v == nil {
		return nil, e
	}
	return &accelGroup{v}, e
}

func unwrapAccelGroup(v gtki.AccelGroup) *gtk.AccelGroup {
	if v == nil {
		return nil
	}
	return v.(*accelGroup).AccelGroup
}
