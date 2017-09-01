package gtki

type Entry interface {
	Widget
	Editable

	GetText() (string, error)
	SetHasFrame(bool)
	SetText(string)
	SetVisibility(bool)
	SetWidthChars(int)
}

func AssertEntry(_ Entry) {}
