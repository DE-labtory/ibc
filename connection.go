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
	GetState() ConnectionState
	GetCounterPartyConnectionIdentifer() Identifier
	GetCounterPartyPrefix() CommitmentPrefix
	GetClientIdentifer() Identifier
	GetCounterPartyClientIdentifer() Identifier
	GetVersion() []string
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

func (ce *ConnectionEndImpl) GetState() ConnectionState {
	return ce.state
}

func (ce *ConnectionEndImpl) GetCounterPartyPrefix() CommitmentPrefix {
	return ce.counterPartyPrefix
}

func (ce *ConnectionEndImpl) GetCounterPartyConnectionIdentifer() Identifier {
	return ce.counterPartyConnectionIdentifer
}

func (ce *ConnectionEndImpl) GetCounterPartyClientIdentifer() Identifier {
	return ce.counterPartyClientIdentifer
}

func (ce *ConnectionEndImpl) GetClientIdentifer() Identifier {
	return ce.clientIdentifier
}

func (ce *ConnectionEndImpl) GetVersion() []string {
	return ce.version
}
