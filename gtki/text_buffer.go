package gtki

import "github.com/twstrike/gotk3adapter/glibi"

type TextBuffer interface {
	glibi.Object

	ApplyTagByName(string, TextIter, TextIter)
	GetCharCount() int
	GetEndIter() TextIter
	GetIterAtOffset(int) TextIter
	GetStartIter() TextIter
	Insert(TextIter, string)
}

func AssertTextBuffer(_ TextBuffer) {}
