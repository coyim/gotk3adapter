package gliba

import "github.com/twstrike/gotk3adapter/glibi"

func init() {
	glibi.AssertGlib(&RealGlib{})
	glibi.AssertObject(&object{})
	glibi.AssertSignal(&signal{})
}
