package pango_mock

import "github.com/twstrike/gotk3adapter/pangoi"

type Mock struct{}

func (*Mock) AsFontDescription(v interface{}) pangoi.FontDescription {
	return nil
}
