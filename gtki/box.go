package gtki

type Box interface {
	Container
	Orientable

	PackEnd(Widget, bool, bool, uint)
	PackStart(Widget, bool, bool, uint)
	SetChildPacking(Widget, bool, bool, uint, PackType)
	SetCenterWidget(Widget)
	GetCenterWidget() Widget
	SetFocusVAdjustment(Adjustment)
}

func AssertBox(_ Box) {}
