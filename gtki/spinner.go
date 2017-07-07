package gtki

import (
	. "github.com/gotk3/gotk3/gtk"
	. "github.com/twstrike/gotk3adapter/gtki"
)

// Spinner is an interface of Gtk.Spinner
type Spinner interface {
	Widget

	Start()
	Stop()
}

// AssertSpinner asserts
func AssertSpinner(_ Spinner) {}
