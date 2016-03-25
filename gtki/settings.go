package gtki

import "github.com/twstrike/gotk3adapter/glibi"

type Settings interface {
	glibi.Object

	GetProperty(string) (interface{}, error)
}

func AssertSettings(_ Settings) {}
