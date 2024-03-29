package gtk_mock

import (
	"github.com/coyim/gotk3adapter/gtki"
)

type MockAssistant struct {
	MockWindow
}

func (a *MockAssistant) Commit() {
}

func (a *MockAssistant) NextPage() {
}

func (a *MockAssistant) PreviousPage() {
}

func (a *MockAssistant) SetCurrentPage(pageNum int) {
}

func (a *MockAssistant) GetCurrentPage() int {
	return 0
}

func (a *MockAssistant) GetNthPage(pageNum int) (gtki.Widget, error) {
	return nil, nil
}

func (a *MockAssistant) AppendPage(page gtki.Widget) int {
	return 0
}

func (a *MockAssistant) SetPageType(page gtki.Widget, ptype gtki.AssistantPageType) {
}

func (a *MockAssistant) GetPageType(page gtki.Widget) gtki.AssistantPageType {
	return gtki.ASSISTANT_PAGE_SUMMARY
}

func (a *MockAssistant) SetPageComplete(page gtki.Widget, complete bool) {
}

func (a *MockAssistant) GetPageComplete(page gtki.Widget) bool {
	return true
}

func (a *MockAssistant) AddActionWidget(child gtki.Widget) {
}

func (a *MockAssistant) RemoveActionWidget(child gtki.Widget) {
}

func (a *MockAssistant) UpdateButtonsState() {
}

func (a *MockAssistant) SetPageTitle(page gtki.Widget, title string) {
}

func (a *MockAssistant) GetButtons() []gtki.Button {
	return nil
}

func (a *MockAssistant) GetButtonSizeGroup() (gtki.SizeGroup, error) {
	return nil, nil
}

func (a *MockAssistant) GetHeaderBar() (gtki.HeaderBar, error) {
	return nil, nil
}

func (a *MockAssistant) GetSidebar() (gtki.Box, error) {
	return nil, nil
}

func (a *MockAssistant) GetNotebook() (gtki.Notebook, error) {
	return nil, nil
}

func (a *MockAssistant) HideBottomActionArea() {
}
