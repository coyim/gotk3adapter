package gtka

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gliba"
)

func init() {
	gliba.AddWrapper(WrapLocal)

	gliba.AddUnwrapper(UnwrapLocal)
}

func Wrap(o interface{}) interface{} {
	v1, ok := WrapLocal(o)
	if !ok {
		panic(fmt.Sprintf("Unrecognized type of object: %#v", o))
	}
	return v1
}

func Unwrap(o interface{}) interface{} {
	v1, ok := UnwrapLocal(o)
	if !ok {
		panic(fmt.Sprintf("Unrecognized type of object: %#v", o))
	}
	return v1
}

func WrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *gtk.AboutDialog:
		return wrapAboutDialogSimple(oo), true
	case *gtk.AccelGroup:
		return wrapAccelGroupSimple(oo), true
	case *gtk.Adjustment:
		return wrapAdjustmentSimple(oo), true
	case *gtk.ApplicationWindow:
		return wrapApplicationWindowSimple(oo), true
	case *gtk.Bin:
		return wrapBinSimple(oo), true
	case *gtk.Box:
		return wrapBoxSimple(oo), true
	case *gtk.Button:
		return wrapButtonSimple(oo), true
	case *gtk.CellRenderer:
		return wrapCellRendererSimple(oo), true
	case *gtk.CellRendererText:
		return wrapCellRendererTextSimple(oo), true
	case *gtk.CellRendererToggle:
		return wrapCellRendererToggleSimple(oo), true
	case *gtk.CheckButton:
		return wrapCheckButtonSimple(oo), true
	case *gtk.CheckMenuItem:
		return wrapCheckMenuItemSimple(oo), true
	case *gtk.ComboBox:
		return wrapComboBoxSimple(oo), true
	case *gtk.ComboBoxText:
		return wrapComboBoxTextSimple(oo), true
	case *gtk.Container:
		return wrapContainerSimple(oo), true
	case *gtk.Dialog:
		return wrapDialogSimple(oo), true
	case *gtk.Entry:
		return wrapEntrySimple(oo), true
	case *gtk.FileChooserDialog:
		return wrapFileChooserDialogSimple(oo), true
	case *gtk.Grid:
		return wrapGridSimple(oo), true
	case *gtk.HeaderBar:
		return wrapHeaderBarSimple(oo), true
	case *gtk.InfoBar:
		return wrapInfoBarSimple(oo), true
	case *gtk.Label:
		return wrapLabelSimple(oo), true
	case *gtk.ListStore:
		return wrapListStoreSimple(oo), true
	case *gtk.Menu:
		return wrapMenuSimple(oo), true
	case *gtk.MenuBar:
		return wrapMenuBarSimple(oo), true
	case *gtk.MenuItem:
		return wrapMenuItemSimple(oo), true
	case *gtk.MenuShell:
		return wrapMenuShellSimple(oo), true
	case *gtk.MessageDialog:
		return wrapMessageDialogSimple(oo), true
	case *gtk.Notebook:
		return wrapNotebookSimple(oo), true
	case *gtk.Revealer:
		return wrapRevealerSimple(oo), true
	case *gtk.ScrolledWindow:
		return wrapScrolledWindowSimple(oo), true
	case *gtk.SeparatorMenuItem:
		return wrapSeparatorMenuItemSimple(oo), true
	case *gtk.TextBuffer:
		return wrapTextBufferSimple(oo), true
	case *gtk.TextTag:
		return wrapTextTagSimple(oo), true
	case *gtk.TextTagTable:
		return wrapTextTagTableSimple(oo), true
	case *gtk.TextView:
		return wrapTextViewSimple(oo), true
	case *gtk.ToggleButton:
		return wrapToggleButtonSimple(oo), true
	case *gtk.TreePath:
		return wrapTreePathSimple(oo), true
	case *gtk.TreeSelection:
		return wrapTreeSelectionSimple(oo), true
	case *gtk.TreeStore:
		return wrapTreeStoreSimple(oo), true
	case *gtk.TreeView:
		return wrapTreeViewSimple(oo), true
	case *gtk.TreeViewColumn:
		return wrapTreeViewColumnSimple(oo), true
	case *gtk.Widget:
		return wrapWidgetSimple(oo), true
	case *gtk.Window:
		return wrapWindowSimple(oo), true
	default:
		return nil, false
	}
}

func UnwrapLocal(o interface{}) (interface{}, bool) {
	switch oo := o.(type) {
	case *aboutDialog:
		return unwrapAboutDialog(oo), true
	case *accelGroup:
		return unwrapAccelGroup(oo), true
	case *adjustment:
		return unwrapAdjustment(oo), true
	case *applicationWindow:
		return unwrapApplicationWindow(oo), true
	case *bin:
		return unwrapBin(oo), true
	case *box:
		return unwrapBox(oo), true
	case *button:
		return unwrapButton(oo), true
	case *cellRenderer:
		return unwrapCellRenderer(oo), true
	case *cellRendererText:
		return unwrapCellRendererText(oo), true
	case *cellRendererToggle:
		return unwrapCellRendererToggle(oo), true
	case *checkButton:
		return unwrapCheckButton(oo), true
	case *checkMenuItem:
		return unwrapCheckMenuItem(oo), true
	case *comboBox:
		return unwrapComboBox(oo), true
	case *comboBoxText:
		return unwrapComboBoxText(oo), true
	case *container:
		return unwrapContainer(oo), true
	case *dialog:
		return unwrapDialog(oo), true
	case *entry:
		return unwrapEntry(oo), true
	case *fileChooserDialog:
		return unwrapFileChooserDialog(oo), true
	case *grid:
		return unwrapGrid(oo), true
	case *headerBar:
		return unwrapHeaderBar(oo), true
	case *infoBar:
		return unwrapInfoBar(oo), true
	case *label:
		return unwrapLabel(oo), true
	case *listStore:
		return unwrapListStore(oo), true
	case *menu:
		return unwrapMenu(oo), true
	case *menuBar:
		return unwrapMenuBar(oo), true
	case *menuItem:
		return unwrapMenuItem(oo), true
	case *menuShell:
		return unwrapMenuShell(oo), true
	case *messageDialog:
		return unwrapMessageDialog(oo), true
	case *notebook:
		return unwrapNotebook(oo), true
	case *revealer:
		return unwrapRevealer(oo), true
	case *scrolledWindow:
		return unwrapScrolledWindow(oo), true
	case *separatorMenuItem:
		return unwrapSeparatorMenuItem(oo), true
	case *textBuffer:
		return unwrapTextBuffer(oo), true
	case *textTag:
		return unwrapTextTag(oo), true
	case *textTagTable:
		return unwrapTextTagTable(oo), true
	case *textView:
		return unwrapTextView(oo), true
	case *toggleButton:
		return unwrapToggleButton(oo), true
	case *treePath:
		return unwrapTreePath(oo), true
	case *treeSelection:
		return unwrapTreeSelection(oo), true
	case *treeStore:
		return unwrapTreeStore(oo), true
	case *treeView:
		return unwrapTreeView(oo), true
	case *treeViewColumn:
		return unwrapTreeViewColumn(oo), true
	case *widget:
		return unwrapWidget(oo), true
	case *window:
		return unwrapWindow(oo), true
	default:
		return nil, false
	}
}
