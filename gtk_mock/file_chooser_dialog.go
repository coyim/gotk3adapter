package gtk_mock

type MockFileChooserDialog struct {
	MockDialog
}

func (*MockFileChooserDialog) GetFilename() string {
	return ""
}

func (*MockFileChooserDialog) SetCurrentName(v1 string) {
}

func (*MockFileChooserDialog) SetDoOverwriteConfirmation(v1 bool) {
}
