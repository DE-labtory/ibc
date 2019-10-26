package spec

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

