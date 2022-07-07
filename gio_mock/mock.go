package gio_mock

import (
	"github.com/coyim/gotk3adapter/gioi"
)

type Mock struct{}

func (*Mock) LoadResource(string) (gioi.Resource, error) {
	return nil, nil
}
func (*Mock) NewResourceFromData([]byte) (gioi.Resource, error) {
	return nil, nil
}
func (*Mock) RegisterResource(gioi.Resource) {
}
func (*Mock) UnregisterResource(gioi.Resource) {
}
