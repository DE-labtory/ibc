package ibc

import (
	"encoding/json"
	"github.com/DE-labtory/ibc/spec"
)

type ProvableStore struct {
	store map[spec.Path]spec.Value
}

func (ps *ProvableStore) get(path spec.Path) spec.Value {
	if value, ok := ps.store[path]; ok {
		return value
	}
	return nil
}

func (ps *ProvableStore) set(path spec.Path, value interface{}) {
	m, err := json.Marshal(value)
	if err != nil {
		return
	}

	ps.store[path] = m

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

func (ps *PrivateStore) set(path spec.Path, value interface{}) {
	m, err := json.Marshal(value)
	if err != nil {
		return
	}

	ps.store[path] = m
}

func (ps *PrivateStore) delete(path spec.Path) {
	delete(ps.store, path)
}
