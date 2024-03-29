package glibi

type Glib interface {
	IdleAdd(interface{}) SourceHandle
	TimeoutAdd(uint, interface{}) SourceHandle
	TimeoutSecondsAdd(uint, interface{}) SourceHandle
	InitI18n(string, string)
	Local(string) string
	MainDepth() int

	SettingsNew(string) Settings
	SettingsNewWithPath(string, string) Settings
	SettingsNewWithBackend(string, SettingsBackend) Settings
	SettingsNewWithBackendAndPath(string, SettingsBackend, string) Settings
	SettingsNewFull(SettingsSchema, SettingsBackend, string) Settings
	SettingsSync()

	SettingsBackendGetDefault() SettingsBackend
	KeyfileSettingsBackendNew(string, string, string) SettingsBackend
	MemorySettingsBackendNew() SettingsBackend
	NullSettingsBackendNew() SettingsBackend

	SettingsSchemaSourceGetDefault() SettingsSchemaSource
	SettingsSchemaSourceNewFromDirectory(string, SettingsSchemaSource, bool) SettingsSchemaSource

	SignalNew(string) (Signal, error)

	MenuNew() Menu
	MenuItemNew(label, detailed_action string) MenuItem
	MenuItemNewSection(label string, section MenuModel) MenuItem
	MenuItemNewSubmenu(label string, submenu MenuModel) MenuItem
	MenuItemNewFromModel(model MenuModel, index int) MenuItem

	ActionNameIsValid(actionName string) bool
	SimpleActionNew(name string, parameterType VariantType) SimpleAction
	SimpleActionNewStateful(name string, parameterType VariantType, state Variant) SimpleAction
	PropertyActionNew(name string, object Object, propertyName string) PropertyAction

	SetFinalizerStrategy(func(func()))

	MarkupEscapeText(string) string
} // end of Glib

func AssertGlib(_ Glib) {}
