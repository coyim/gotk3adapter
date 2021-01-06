package gtki

type Entry interface {
	Widget
	Editable

	GetText() (string, error)
	SetHasFrame(bool)
	SetText(string)
	SetVisibility(bool)
	SetWidthChars(int)
	GetAlignment() float32
	SetAlignment(float32)
	SetCompletion(EntryCompletion)
}

func AssertEntry(_ Entry) {}
