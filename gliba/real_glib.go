package gliba

import "github.com/gotk3/gotk3/glib"

import "github.com/twstrike/gotk3adapter/glibi"

type RealGlib struct{}

var Real = &RealGlib{}

func (*RealGlib) IdleAdd(f interface{}, args ...interface{}) (glibi.SourceHandle, error) {
	res, err := glib.IdleAdd(f, args...)
	return glibi.SourceHandle(res), err
}

func (*RealGlib) InitI18n(domain string, dir string) {
	glib.InitI18n(domain, dir)
}

func (*RealGlib) SignalNew(s string) (glibi.Signal, error) {
	return wrapSignal(glib.SignalNew(s))
}
