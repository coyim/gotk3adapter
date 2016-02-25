package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type application struct {
	*gtk.Application
}

func wrapApplication(v *gtk.Application, e error) (*application, error) {
	if v == nil {
		return nil, e
	}
	return &application{v}, e
}

func unwrapApplication(v gtki.Application) *gtk.Application {
	if v == nil {
		return nil
	}
	return v.(*application).Application
}
