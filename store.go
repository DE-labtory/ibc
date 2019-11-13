package ibc

import "github.com/DE-labtory/ibc/spec"

type ProvableStore struct {
	store map[spec.Path]spec.Value
}

func (ps *ProvableStore) get(path spec.Path) spec.Value {
	if value, ok := ps.store[path]; ok {
		return value
	}
	return nil
}

func (ps *ProvableStore) set(path spec.Path, value spec.Value) {
	ps.store[path] = value
}

func (ps *ProvableStore) delete(path spec.Path) {
	delete(ps.store, path)
}

type PrivateStore struct {
	store map[spec.Path]spec.Value
}

func (ps *PrivateStore) get(path spec.Path) spec.Value {
	if value, ok := ps.store[path]; ok {
		return value
	}
	return nil
}

func (ps *PrivateStore) set(path spec.Path, value spec.Value) {
	ps.store[path] = value
}

func (ps *PrivateStore) delete(path spec.Path) {
	delete(ps.store, path)
}
