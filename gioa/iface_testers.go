package gioa

import "github.com/coyim/gotk3adapter/gioi"

func init() {
	gioi.AssertGio(&RealGio{})
	gioi.AssertResource(&resource{})
}
