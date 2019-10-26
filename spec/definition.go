package spec

import "fmt"

type Identifier string

// TODO : MUST REDEFINE IT!
type Value interface{}

type Path string

// todo: implement
// Paths MUST contain only identifiers, constant alphanumeric strings, and the separator "/"
func (p Path) isValid() bool {
	return false
}

func ClientStatePath(id Identifier) Path {
	return Path(fmt.Sprintf("clients/%s/state", id))
}

type Datagram struct {
}

type Packet struct {
	sequence      int
	timeoutHeight int
	sourcePort    Identifier
	sourceChannel Identifier
	destPort      Identifier
	destChannel   Identifier
	data          []byte
}

type OpaquePacket interface{}
