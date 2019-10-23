package ibc

// ConnectionState provides current state of the connection end.
type ConnectionState int

const (
	// CONNINIT means state machine is initialized.
	CONNINIT ConnectionState = 0

	// CONNOPENTRY means state machine is trying to open connection.
	CONNOPENTRY

	// CONNOPEN means state machine opens connection.
	CONNOPEN
)

// ConnectionEnd is a interface for connection end.
type ConnectionEnd interface {
	getState() ConnectionState
	getCounterPartyConnectionIdentifer() Identifier
	getCounterPartyPrefix() CommitmentPrefix
	getClientIdentifer() Identifier
	getCounterPartyClientIdentifer() Identifier
	getVersion() []string
}

// ConnectionEndImpl is a temporarily implementation ConnectionEnd.
type ConnectionEndImpl struct {
	state                           ConnectionState
	counterPartyConnectionIdentifer Identifier
	counterPartyPrefix              CommitmentPrefix
	clientIdentifier                Identifier
	counterPartyClientIdentifer     Identifier
	version                         []string
}

func (ce *ConnectionEndImpl) getState() ConnectionState {
	return ce.state
}

func (ce *ConnectionEndImpl) getCounterPartyPrefix() CommitmentPrefix {
	return ce.counterPartyPrefix
}

func (ce *ConnectionEndImpl) getCounterPartyConnectionIdentifer() Identifier {
	return ce.counterPartyConnectionIdentifer
}

func (ce *ConnectionEndImpl) getCounterPartyClientIdentifer() Identifier {
	return ce.counterPartyClientIdentifer
}

func (ce *ConnectionEndImpl) getClientIdentifer() Identifier {
	return ce.clientIdentifier
}

func (ce *ConnectionEndImpl) getVersion() []string {
	return ce.version
}
