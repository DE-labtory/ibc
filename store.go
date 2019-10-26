package ibc

import "github.com/DE-labtory/ibc/spec"

type ProvableStore struct {
}

func (*ProvableStore) get(path spec.Path) spec.Value {
	panic("implement me")
}

func (*ProvableStore) set(path spec.Path, value spec.Value) {
	panic("implement me")
}

func (*ProvableStore) delete(path spec.Path) {
	panic("implement me")
}

type PrivateStore struct {
}

func (*PrivateStore) get(path spec.Path) spec.Value {
	panic("implement me")
}

func (*PrivateStore) set(path spec.Path, value spec.Value) {
	panic("implement me")
}

func (*PrivateStore) delete(path spec.Path) {
	panic("implement me")
}
