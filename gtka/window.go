package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gdka"
	"github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/gotk3adapter/gtki"
)

type window struct {
	*bin
	*gtk.Window
}

func wrapWindowSimple(v *gtk.Window) *window {
	if v == nil {
		return nil
	}
	return &window{wrapBinSimple(&v.Bin), v}
}

func wrapWindow(v *gtk.Window, e error) (*window, error) {
	return wrapWindowSimple(v), e
}

func unwrapWindow(v gtki.Window) *gtk.Window {
	if v == nil {
		return nil
	}
	return v.(*window).Window
}

func (v *window) AddAccelGroup(v2 gtki.AccelGroup) {
	v.Window.AddAccelGroup(unwrapAccelGroup(v2))
}

func (v *window) SetApplication(v2 gtki.Application) {
	v.Window.SetApplication(unwrapApplication(v2))
}

func (v *window) SetIcon(v2 gdki.Pixbuf) {
	v.Window.SetIcon(gdka.UnwrapPixbuf(v2))
}

func (v *window) SetTitlebar(v2 gtki.Widget) {
	v.Window.SetTitlebar(unwrapWidget(v2))
}

func (v *window) SetTransientFor(v2 gtki.Window) {
	v.Window.SetTransientFor(unwrapWindow(v2))
}
