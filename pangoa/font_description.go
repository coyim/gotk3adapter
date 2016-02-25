package pangoa

import (
	"github.com/gotk3/gotk3/pango"
	"github.com/twstrike/gotk3adapter/pangoi"
)

type fontDescription struct {
	*pango.FontDescription
}

func wrapFontDescription(p *pango.FontDescription) pangoi.FontDescription {
	return &fontDescription{p}
}
