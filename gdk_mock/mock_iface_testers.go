package gdk_mock

import "github.com/twstrike/gotk3adapter/gdki"

func init() {
	gdki.AssertGdk(&Mock{})
	gdki.AssertEvent(&MockEvent{})
	gdki.AssertEventButton(&MockEventButton{})
	gdki.AssertPixbuf(&MockPixbuf{})
	gdki.AssertPixbufLoader(&MockPixbufLoader{})
	gdki.AssertScreen(&MockScreen{})
	gdki.AssertWindow(&MockWindow{})
}
