package gtk_mock

import "github.com/twstrike/gotk3adapter/gtki"

type MockImage struct {
	MockWidget
}

func (a *MockImage) SetFromIconName(v1 string, v2 gtki.IconSize) {
}
