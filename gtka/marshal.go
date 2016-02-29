package gtka

import "github.com/twstrike/gotk3adapter/gtki"
import "github.com/gotk3/gotk3/gtk"

func marshalAboutDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapAboutDialog(v.(gtki.AboutDialog))
}

func unmarshalAboutDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapAboutDialogSimple(v.(*gtk.AboutDialog))
}

func marshalAccelGroup(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapAccelGroup(v.(gtki.AccelGroup))
}

func unmarshalAccelGroup(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapAccelGroupSimple(v.(*gtk.AccelGroup))
}

func marshalAdjustment(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapAdjustment(v.(gtki.Adjustment))
}

func unmarshalAdjustment(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapAdjustmentSimple(v.(*gtk.Adjustment))
}

func marshalApplication(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapApplication(v.(gtki.Application))
}

func unmarshalApplication(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapApplicationSimple(v.(*gtk.Application))
}

func marshalApplicationWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapApplicationWindow(v.(gtki.ApplicationWindow))
}

func unmarshalApplicationWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapApplicationWindowSimple(v.(*gtk.ApplicationWindow))
}

func marshalBox(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapBox(v.(gtki.Box))
}

func unmarshalBox(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapBoxSimple(v.(*gtk.Box))
}

func marshalBuilder(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapBuilder(v.(gtki.Builder))
}

func unmarshalBuilder(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapBuilderSimple(v.(*gtk.Builder))
}

func marshalButton(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapButton(v.(gtki.Button))
}

func unmarshalButton(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapButtonSimple(v.(*gtk.Button))
}

func marshalCellRendererText(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapCellRendererText(v.(gtki.CellRendererText))
}

func unmarshalCellRendererText(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapCellRendererTextSimple(v.(*gtk.CellRendererText))
}

func marshalCellRendererToggle(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapCellRendererToggle(v.(gtki.CellRendererToggle))
}

func unmarshalCellRendererToggle(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapCellRendererToggleSimple(v.(*gtk.CellRendererToggle))
}

func marshalCheckButton(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapCheckButton(v.(gtki.CheckButton))
}

func unmarshalCheckButton(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapCheckButtonSimple(v.(*gtk.CheckButton))
}

func marshalCheckMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapCheckMenuItem(v.(gtki.CheckMenuItem))
}

func unmarshalCheckMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapCheckMenuItemSimple(v.(*gtk.CheckMenuItem))
}

func marshalComboBox(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapComboBox(v.(gtki.ComboBox))
}

func unmarshalComboBox(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapComboBoxSimple(v.(*gtk.ComboBox))
}

func marshalComboBoxText(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapComboBoxText(v.(gtki.ComboBoxText))
}

func unmarshalComboBoxText(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapComboBoxTextSimple(v.(*gtk.ComboBoxText))
}

func marshalCssProvider(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapCssProvider(v.(gtki.CssProvider))
}

func unmarshalCssProvider(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapCssProviderSimple(v.(*gtk.CssProvider))
}

func marshalDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapDialog(v.(gtki.Dialog))
}

func unmarshalDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapDialogSimple(v.(*gtk.Dialog))
}

func marshalEntry(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapEntry(v.(gtki.Entry))
}

func unmarshalEntry(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapEntrySimple(v.(*gtk.Entry))
}

func marshalFileChooserDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapFileChooserDialog(v.(gtki.FileChooserDialog))
}

func unmarshalFileChooserDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapFileChooserDialogSimple(v.(*gtk.FileChooserDialog))
}

func marshalGrid(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapGrid(v.(gtki.Grid))
}

func unmarshalGrid(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapGridSimple(v.(*gtk.Grid))
}

func marshalInfoBar(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapInfoBar(v.(gtki.InfoBar))
}

func unmarshalInfoBar(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapInfoBarSimple(v.(*gtk.InfoBar))
}

func marshalLabel(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapLabel(v.(gtki.Label))
}

func unmarshalLabel(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapLabelSimple(v.(*gtk.Label))
}

func marshalListStore(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapListStore(v.(gtki.ListStore))
}

func unmarshalListStore(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapListStoreSimple(v.(*gtk.ListStore))
}

func marshalMenu(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapMenu(v.(gtki.Menu))
}

func unmarshalMenu(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapMenuSimple(v.(*gtk.Menu))
}

func marshalMenuBar(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapMenuBar(v.(gtki.MenuBar))
}

func unmarshalMenuBar(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapMenuBarSimple(v.(*gtk.MenuBar))
}

func marshalMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapMenuItem(v.(gtki.MenuItem))
}

func unmarshalMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapMenuItemSimple(v.(*gtk.MenuItem))
}

func marshalMessageDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapMessageDialog(v.(gtki.MessageDialog))
}

func unmarshalMessageDialog(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapMessageDialogSimple(v.(*gtk.MessageDialog))
}

func marshalNotebook(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapNotebook(v.(gtki.Notebook))
}

func unmarshalNotebook(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapNotebookSimple(v.(*gtk.Notebook))
}

func marshalRevealer(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapRevealer(v.(gtki.Revealer))
}

func unmarshalRevealer(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapRevealerSimple(v.(*gtk.Revealer))
}

func marshalScrolledWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapScrolledWindow(v.(gtki.ScrolledWindow))
}

func unmarshalScrolledWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapScrolledWindowSimple(v.(*gtk.ScrolledWindow))
}

func marshalSeparatorMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapSeparatorMenuItem(v.(gtki.SeparatorMenuItem))
}

func unmarshalSeparatorMenuItem(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapSeparatorMenuItemSimple(v.(*gtk.SeparatorMenuItem))
}

func marshalTextBuffer(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTextBuffer(v.(gtki.TextBuffer))
}

func unmarshalTextBuffer(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTextBufferSimple(v.(*gtk.TextBuffer))
}

func marshalTextIter(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTextIter(v.(gtki.TextIter))
}

func unmarshalTextIter(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTextIterSimple(v.(*gtk.TextIter))
}

func marshalTextTag(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTextTag(v.(gtki.TextTag))
}

func unmarshalTextTag(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTextTagSimple(v.(*gtk.TextTag))
}

func marshalTextTagTable(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTextTagTable(v.(gtki.TextTagTable))
}

func unmarshalTextTagTable(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTextTagTableSimple(v.(*gtk.TextTagTable))
}

func marshalTextView(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTextView(v.(gtki.TextView))
}

func unmarshalTextView(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTextViewSimple(v.(*gtk.TextView))
}

func marshalTreeIter(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTreeIter(v.(gtki.TreeIter))
}

func unmarshalTreeIter(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTreeIterSimple(v.(*gtk.TreeIter))
}

func marshalTreePath(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTreePath(v.(gtki.TreePath))
}

func unmarshalTreePath(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTreePathSimple(v.(*gtk.TreePath))
}

func marshalTreeSelection(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTreeSelection(v.(gtki.TreeSelection))
}

func unmarshalTreeSelection(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTreeSelectionSimple(v.(*gtk.TreeSelection))
}

func marshalTreeStore(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTreeStore(v.(gtki.TreeStore))
}

func unmarshalTreeStore(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTreeStoreSimple(v.(*gtk.TreeStore))
}

func marshalTreeView(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapTreeView(v.(gtki.TreeView))
}

func unmarshalTreeView(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapTreeViewSimple(v.(*gtk.TreeView))
}

func marshalWidget(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapWidget(v.(gtki.Widget))
}

func unmarshalWidget(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapWidgetSimple(v.(*gtk.Widget))
}

func marshalWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return unwrapWindow(v.(gtki.Window))
}

func unmarshalWindow(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return wrapWindowSimple(v.(*gtk.Window))
}
