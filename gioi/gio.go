package gioi

type Gio interface {
	LoadResource(string) (Resource, error)
	NewResourceFromData([]byte) (Resource, error)
	RegisterResource(Resource)
	UnregisterResource(Resource)
}

func AssertGio(_ Gio) {}
