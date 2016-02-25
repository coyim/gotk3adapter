package gdka

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/gotk3adapter/gdki"
)

type screen struct {
	*gdk.Screen
}

func wrapScreen(v *gdk.Screen, e error) (*screen, error) {
	if v == nil {
		return nil, e
	}
	return &screen{v}, e
}

func UnwrapScreen(v gdki.Screen) *gdk.Screen {
	if v == nil {
		return nil
	}
	return v.(*screen).Screen
}
