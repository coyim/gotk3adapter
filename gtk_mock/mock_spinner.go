package gtk_mock

import . "github.com/twstrike/gotk3adapter/gtki"

type MockSpinner struct {
	MockWidget
}

func (*MockSpinner) Start() {
}

func (*MockSpinner) Stop() {
}
