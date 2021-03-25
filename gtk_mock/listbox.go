package gtk_mock

import "gotk3adapter/gtki"

type MockListBox struct {
	MockContainer
}

func (*MockListBox) SelectRow(gtki.ListBoxRow) {
}

func (*MockListBox) GetRowAtIndex(int) gtki.ListBoxRow {
	return nil
}
