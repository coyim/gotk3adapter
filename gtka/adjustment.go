package gtka

import (
	"github.com/coyim/gotk3adapter/gliba"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/gotk3/gotk3/gtk"
)

type adjustment struct {
	*gliba.Object
	internal *gtk.Adjustment
}

func WrapAdjustmentSimple(v *gtk.Adjustment) gtki.Adjustment {
	if v == nil {
		return nil
	}
	return &adjustment{gliba.WrapObjectSimple(v.Object), v}
}

func WrapAdjustment(v *gtk.Adjustment, e error) (gtki.Adjustment, error) {
	return WrapAdjustmentSimple(v), e
}

func UnwrapAdjustment(v gtki.Adjustment) *gtk.Adjustment {
	if v == nil {
		return nil
	}
	return v.(*adjustment).internal
}

func (v *adjustment) GetLower() float64 {
	return v.internal.GetLower()
}

func (v *adjustment) GetPageSize() float64 {
	return v.internal.GetPageSize()
}

func (v *adjustment) GetUpper() float64 {
	return v.internal.GetUpper()
}

func (v *adjustment) SetValue(v1 float64) {
	v.internal.SetValue(v1)
}
