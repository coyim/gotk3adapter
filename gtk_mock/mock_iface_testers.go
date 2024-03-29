package gtk_mock

import "github.com/coyim/gotk3adapter/gtki"

func init() {
	gtki.AssertGtk(&Mock{})
	gtki.AssertAboutDialog(&MockAboutDialog{})
	gtki.AssertAccelGroup(&MockAccelGroup{})
	gtki.AssertAdjustment(&MockAdjustment{})
	gtki.AssertAllocation(&MockAllocation{})
	gtki.AssertApplication(&MockApplication{})
	gtki.AssertApplicationWindow(&MockApplicationWindow{})
	gtki.AssertAssistant(&MockAssistant{})
	gtki.AssertBox(&MockBox{})
	gtki.AssertBuilder(&MockBuilder{})
	gtki.AssertButton(&MockButton{})
	gtki.AssertCellRenderer(&MockCellRenderer{})
	gtki.AssertCellRendererText(&MockCellRendererText{})
	gtki.AssertCellRendererToggle(&MockCellRendererToggle{})
	gtki.AssertCheckButton(&MockCheckButton{})
	gtki.AssertCheckMenuItem(&MockCheckMenuItem{})
	gtki.AssertComboBox(&MockComboBox{})
	gtki.AssertComboBoxText(&MockComboBoxText{})
	gtki.AssertCssProvider(&MockCssProvider{})
	gtki.AssertDialog(&MockDialog{})
	gtki.AssertEntry(&MockEntry{})
	gtki.AssertEventBox(&MockEventBox{})
	gtki.AssertExpander(&MockExpander{})
	gtki.AssertButtonBox(&MockButtonBox{})
	gtki.AssertFileChooserDialog(&MockFileChooserDialog{})
	gtki.AssertGrid(&MockGrid{})
	gtki.AssertHeaderBar(&MockHeaderBar{})
	gtki.AssertImage(&MockImage{})
	gtki.AssertInfoBar(&MockInfoBar{})
	gtki.AssertLabel(&MockLabel{})
	gtki.AssertListStore(&MockListStore{})
	gtki.AssertMenuBar(&MockMenuBar{})
	gtki.AssertMenuItem(&MockMenuItem{})
	gtki.AssertMenu(&MockMenu{})
	gtki.AssertMessageDialog(&MockMessageDialog{})
	gtki.AssertNotebook(&MockNotebook{})
	gtki.AssertProgressBar(&MockProgressBar{})
	gtki.AssertRevealer(&MockRevealer{})
	gtki.AssertScrolledWindow(&MockScrolledWindow{})
	gtki.AssertSearchBar(&MockSearchBar{})
	gtki.AssertSearchEntry(&MockSearchEntry{})
	gtki.AssertSeparatorMenuItem(&MockSeparatorMenuItem{})
	gtki.AssertSettings(&MockSettings{})
	gtki.AssertSpinButton(&MockSpinButton{})
	gtki.AssertSpinner(&MockSpinner{})
	gtki.AssertStyleContext(&MockStyleContext{})
	gtki.AssertTextBuffer(&MockTextBuffer{})
	gtki.AssertTextIter(&MockTextIter{})
	gtki.AssertTextMark(&MockTextMark{})
	gtki.AssertTextTagTable(&MockTextTagTable{})
	gtki.AssertTextTag(&MockTextTag{})
	gtki.AssertTextView(&MockTextView{})
	gtki.AssertTreeIter(&MockTreeIter{})
	gtki.AssertTreePath(&MockTreePath{})
	gtki.AssertTreeSelection(&MockTreeSelection{})
	gtki.AssertTreeStore(&MockTreeStore{})
	gtki.AssertTreeView(&MockTreeView{})
	gtki.AssertTreeViewColumn(&MockTreeViewColumn{})
	gtki.AssertWidget(&MockWidget{})
	gtki.AssertWindow(&MockWindow{})
	gtki.AssertIconTheme(&MockIconTheme{})
}
