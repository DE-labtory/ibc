package ibc

type KVStore interface {
	get(path Path) Value
	set(path Path, value Value)
	delete(path Path)
}

type ProvableStore struct {
}

func (*ProvableStore) get(path Path) Value {
	panic("implement me")
}

func (*ProvableStore) set(path Path, value Value) {
	panic("implement me")
}

func (*ProvableStore) delete(path Path) {
	panic("implement me")
}

type PrivateStore struct {
}

func (*PrivateStore) get(path Path) Value {
	panic("implement me")
}

func (*PrivateStore) set(path Path, value Value) {
	panic("implement me")
}

func (*PrivateStore) delete(path Path) {
	panic("implement me")
}
