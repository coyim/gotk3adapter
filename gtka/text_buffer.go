package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type textBuffer struct {
	*gtk.TextBuffer
}

func wrapTextBuffer(v *gtk.TextBuffer, e error) (*textBuffer, error) {
	if v == nil {
		return nil, e
	}
	return &textBuffer{v}, e
}

func unwrapTextBuffer(v gtki.TextBuffer) *gtk.TextBuffer {
	if v == nil {
		return nil
	}
	return v.(*textBuffer).TextBuffer
}
