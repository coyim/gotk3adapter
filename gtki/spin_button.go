package gtki

import (
	. "github.com/gotk3/gotk3/gtk"
	. "github.com/twstrike/gotk3adapter/gtki"
)

type SpinButton interface {
	Entry

	GetValueAsInt() int
	SetValue(float64)
	GetValue() float64
	GetAdjustment() Adjustment
	SetRange(float64, float64)
	SetIncrements(float64, float64)
}

func AssertSpinButton(_ SpinButton) {}
