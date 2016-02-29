package gliba

import "github.com/gotk3/gotk3/glib"

type Application struct {
	*Object
	*glib.Application
}

func WrapApplicationSimple(v *glib.Application) *Application {
	if v == nil {
		return nil
	}
	return &Application{WrapObjectSimple(v.Object), v}
}
