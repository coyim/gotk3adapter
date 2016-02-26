package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type accelGroup struct {
	*gliba.Object
	*gtk.AccelGroup
}

func wrapAccelGroupSimple(v *gtk.AccelGroup) *accelGroup {
	if v == nil {
		return nil
	}
	return &accelGroup{gliba.WrapObjectSimple(&v.Object), v}
}

func wrapAccelGroup(v *gtk.AccelGroup, e error) (*accelGroup, error) {
	return wrapAccelGroupSimple(v), e
}

func unwrapAccelGroup(v gtki.AccelGroup) *gtk.AccelGroup {
	if v == nil {
		return nil
	}
	return v.(*accelGroup).AccelGroup
}
