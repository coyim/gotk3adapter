package gtki

import "github.com/coyim/gotk3adapter/gdki"

type Allocation interface {
	gdki.Rectangle
	GetY() int
}

func AssertAllocation(_ Allocation) {}
