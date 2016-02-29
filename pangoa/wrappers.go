package pangoa

import (
	"github.com/gotk3/gotk3/pango"
	"github.com/twstrike/gotk3adapter/gliba"
)

func init() {
	gliba.AddWrapper(WrapLocal)

	gliba.AddUnwrapper(UnwrapLocal)
}

func WrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *pango.FontDescription:
		return wrapFontDescriptionSimple(oo), true
	default:
		return nil, false
	}
}

func UnwrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *fontDescription:
		return unwrapFontDescription(oo), true
	default:
		return nil, false
	}
}
