package gliba

import "github.com/gotk3/gotk3/glib"

type signal struct {
	*glib.Signal
}

func wrapSignal(s *glib.Signal, e error) (*signal, error) {
	if e != nil {
		return nil, e
	}
	return &signal{s}, nil
}
