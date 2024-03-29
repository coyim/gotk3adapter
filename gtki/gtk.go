package gtki

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/glibi"
)

type Gtk interface {
	AboutDialogNew() (AboutDialog, error)
	AccelGroupNew() (AccelGroup, error)
	AcceleratorParse(string) (uint, gdki.ModifierType)
	AddProviderForScreen(gdki.Screen, StyleProvider, uint)
	ApplicationNew(string, glibi.ApplicationFlags) (Application, error)
	ApplicationWindowNew(Application) (ApplicationWindow, error)
	AssistantNew() (Assistant, error)
	BuilderNew() (Builder, error)
	BuilderNewFromResource(string) (Builder, error)
	CellRendererTextNew() (CellRendererText, error)
	CheckButtonNew() (CheckButton, error)
	CheckButtonNewWithMnemonic(string) (CheckButton, error)
	CheckMenuItemNewWithMnemonic(string) (CheckMenuItem, error)
	CheckVersion(major, minor, micro uint) error
	ComboBoxNew() (ComboBox, error)
	ComboBoxTextNew() (ComboBoxText, error)
	CssProviderNew() (CssProvider, error)
	CssProviderGetDefault() (CssProvider, error)
	CssProviderGetNamed(string, string) (CssProvider, error)
	EntryNew() (Entry, error)
	EventBoxNew() (EventBox, error)
	ButtonNewWithLabel(string) (Button, error)
	ButtonBoxNew(Orientation) (ButtonBox, error)
	FileChooserDialogNewWith2Buttons(string, Window, FileChooserAction, string, ResponseType, string, ResponseType) (FileChooserDialog, error)
	GetMajorVersion() uint
	GetMinorVersion() uint
	GetMicroVersion() uint
	IconThemeNew() (IconTheme, error)
	IconThemeGetDefault() IconTheme
	IconThemeGetForScreen(gdki.Screen) IconTheme
	ImageNewFromFile(string) (Image, error)
	ImageNewFromResource(string) (Image, error)
	ImageNewFromPixbuf(gdki.Pixbuf) (Image, error)
	ImageNewFromIconName(string, IconSize) (Image, error)
	Init(*[]string)
	InfoBarNew() (InfoBar, error)
	LabelNew(string) (Label, error)
	ListStoreNew(...glibi.Type) (ListStore, error)
	TreeStoreNew(...glibi.Type) (TreeStore, error)
	MenuBarNew() (MenuBar, error)
	MenuItemNew() (MenuItem, error)
	MenuItemNewWithLabel(string) (MenuItem, error)
	MenuItemNewWithMnemonic(string) (MenuItem, error)
	MenuNew() (Menu, error)
	SearchBarNew() (SearchBar, error)
	SearchEntryNew() (SearchEntry, error)
	SeparatorMenuItemNew() (SeparatorMenuItem, error)
	TextBufferNew(TextTagTable) (TextBuffer, error)
	TextTagNew(string) (TextTag, error)
	TextTagTableNew() (TextTagTable, error)
	TextViewNew() (TextView, error)
	TreePathNew() TreePath
	SpinnerNew() (Spinner, error)
	PopoverNew(Widget) (Popover, error)
	BoxNew(Orientation, int) (Box, error)
	WindowSetDefaultIcon(gdki.Pixbuf)
	SettingsGetDefault() (Settings, error)
	SeparatorNew(Orientation) (Separator, error)
	EntryCompletionNew() (EntryCompletion, error)
	WindowNew(WindowType) (Window, error)

	StatusIconNew() (StatusIcon, error)
	StatusIconNewFromFile(filename string) (StatusIcon, error)
	StatusIconNewFromIconName(iconName string) (StatusIcon, error)
	StatusIconNewFromPixbuf(pixbuf gdki.Pixbuf) (StatusIcon, error)

	GetWidgetBuildableName(Widget) (string, error)

	InfoBarSetRevealed(InfoBar, bool)
	InfoBarGetRevealed(InfoBar) bool
}

func AssertGtk(_ Gtk) {}
