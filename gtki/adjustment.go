package gtki

import "github.com/twstrike/gotk3adapter/glibi"

type Adjustment interface {
	glibi.Object

	GetPageSize() float64
	GetUpper() float64
	SetValue(float64)
} // end of Adjustment

func AssertAdjustment(_ Adjustment) {}
