package gtk_mock

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/glib_mock"
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockWidget struct {
	glib_mock.MockObject
}

func (*MockWidget) Map() {
}

func (*MockWidget) SetHExpand(v1 bool) {
}

func (*MockWidget) SetVExpand(v1 bool) {
}

func (*MockWidget) SetSensitive(v1 bool) {
}

func (*MockWidget) IsSensitive() bool {
	return false
}

func (*MockWidget) SetTooltipText(v1 string) {
}

func (*MockWidget) SetVisible(v1 bool) {
}

func (*MockWidget) IsVisible() bool {
	return false
}

func (*MockWidget) SetName(v1 string) {
}

func (*MockWidget) SetNoShowAll(noShow bool) {
}

func (*MockWidget) SetMarginTop(v1 int) {
}

func (*MockWidget) SetMarginBottom(v1 int) {
}

func (*MockWidget) SetSizeRequest(v1, v2 int) {
}

func (*MockWidget) GetAllocatedHeight() int {
	return 0
}

func (*MockWidget) GetName() (string, error) {
	return "", nil
}

func (*MockWidget) GetAllocatedWidth() int {
	return 0
}

func (*MockWidget) GetAllocation() gtki.Allocation {
	return nil
}

func (*MockWidget) GetParent() (gtki.Widget, error) {
	return nil, nil
}

func (*MockWidget) GetParentX() (gtki.Widget, error) {
	return nil, nil
}

func (*MockWidget) GrabFocus() {
}

func (*MockWidget) GrabDefault() {
}

func (*MockWidget) HasFocus() bool {
	return false
}

func (*MockWidget) Hide() {
}

func (*MockWidget) HideOnDelete() {
}

func (*MockWidget) SetCanFocus(v1 bool) {
}

func (*MockWidget) Show() {
}

func (*MockWidget) ShowAll() {
}

func (*MockWidget) GetWindow() (gdki.Window, error) {
	return nil, nil
}

func (*MockWidget) GetStyleContext() (gtki.StyleContext, error) {
	return nil, nil
}

func (*MockWidget) SetHAlign(v2 gtki.Align) {
}

func (*MockWidget) SetVAlign(v2 gtki.Align) {
}

func (*MockWidget) SetOpacity(v2 float64) {
}

func (*MockWidget) Destroy() {
}

func (*MockWidget) TemplateChild(string) (glibi.Object, error) {
	return nil, nil
}
