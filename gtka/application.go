package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gtki"
)

type application struct {
	*gliba.Application
	*gtk.Application
}

func wrapApplicationSimple(v *gtk.Application) *application {
	if v == nil {
		return nil
	}
	return &application{gliba.WrapApplicationSimple(&v.Application), v}
}

func wrapApplication(v *gtk.Application, e error) (*application, error) {
	return wrapApplicationSimple(v), e
}

func unwrapApplication(v gtki.Application) *gtk.Application {
	if v == nil {
		return nil
	}
	return v.(*application).Application
}
