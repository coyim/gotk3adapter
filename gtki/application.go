package gtki

import "github.com/twstrike/gotk3adapter/glibi"

type Application interface {
	glibi.Application
}

func AssertApplication(_ Application) {}
