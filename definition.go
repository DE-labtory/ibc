package ibc

type Identifier string
type Value string

type Path string

// todo: implement
// Paths MUST contain only identifiers, constant alphanumeric strings, and the separator "/"
func (p Path) isValid() bool {
	return false
}

type ConsensusStatue bool
type CommitmentPrefix string

type Datagram struct {
}
