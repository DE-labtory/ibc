package spec

type KVStore interface {
	get(path Path) Value
	set(path Path, value Value)
	delete(path Path)
}