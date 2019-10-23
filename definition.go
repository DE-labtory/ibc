package ibc

type Identifier string
type Value string

type Path string

// todo: implement
// Paths MUST contain only identifiers, constant alphanumeric strings, and the separator "/"
func (p Path) isValid() bool {
	return false
}

type Datagram struct {
}

type Packet struct {
	sequence int
	timeoutHeight int
	sourcePort Identifier
	sourceChannel Identifier
	destPort Identifier
	destChannel Identifier
	data []byte
}

type OpaquePacket interface{}