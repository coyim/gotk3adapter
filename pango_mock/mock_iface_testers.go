package pango_mock

import "github.com/twstrike/gotk3adapter/pangoi"

func init() {
	pangoi.AssertPango(&Mock{})
	pangoi.AssertFontDescription(&MockFontDescription{})
}
